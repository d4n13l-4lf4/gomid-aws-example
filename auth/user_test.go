package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/cucumber/godog"
	"github.com/d4n13l-4lf4/gomid-aws-example/auth"
	lambda "github.com/d4n13l-4lf4/gomid-aws-example/http"
	"github.com/d4n13l-4lf4/gomid-aws-example/internal"
	"github.com/d4n13l-4lf4/gomid-aws-example/mocks/thirdparty/middleware"
	gomid "github.com/d4n13l-4lf4/gomid/middleware"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type (
	AuthenticationTest struct {
		asserter *internal.Asserter
		fn       gomid.Next
		nxt      *middleware.Next
		out      any
		err      error
	}
)

func (t *AuthenticationTest) isAuthenticated(authenticated bool) func(ctx context.Context, user string) error {
	return func(ctx context.Context, user string) error {
		t.asserter = internal.NewAsserter()
		t.nxt = &middleware.Next{}
		t.nxt.EXPECT().
			Execute(ctx, mock.AnythingOfType("*events.APIGatewayProxyRequest")).
			Return("Hello!", nil).
			Maybe()

		authenticatedUsers := lo.Ternary[[]string](authenticated, []string{user}, []string{})

		t.fn = auth.AuthenticateUser(authenticatedUsers)(t.nxt.Execute)
		return nil
	}
}

func (t *AuthenticationTest) submitsRequestForService(ctx context.Context, user string) error {
	in := &events.APIGatewayProxyRequest{
		Headers: map[string]string{
			auth.HeaderUsername: user,
		},
	}
	t.out, t.err = t.fn(ctx, in)
	return nil
}

func (t *AuthenticationTest) submitsWrongRequest(ctx context.Context) error {
	in := 1
	t.out, t.err = t.fn(ctx, in)

	return nil
}

func (t *AuthenticationTest) shouldAllowAccess(ctx context.Context) error {
	assertions := assert.New(t.asserter)

	assertions.Equal(t.out, "Hello!")
	assertions.NoError(t.err)
	t.nxt.AssertExpectations(t.asserter)

	return t.asserter.Error()
}

func (t *AuthenticationTest) shouldDenyAccess(ctx context.Context, msg string, statusCode int) error {
	assertions := assert.New(t.asserter)
	expectedResponse := &lambda.HTTPLambdaResponse[string]{
		Body:       msg,
		StatusCode: statusCode,
	}
	assertions.Equal(t.out, expectedResponse)
	assertions.NoError(t.err)
	t.nxt.AssertExpectations(t.asserter)
	return t.asserter.Error()
}

func InitializeUserAuthenticationTest(sc *godog.ScenarioContext) {
	var authenticationTest AuthenticationTest

	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		authenticationTest = AuthenticationTest{}
		return ctx, nil
	})

	nameRegex := "[A-Za-z0-9\\s]+"
	statusCodeRegex := "\\d{3}"
	sc.Given(fmt.Sprintf("^(%s) is authenticated$", nameRegex), authenticationTest.isAuthenticated(true))
	sc.Given(fmt.Sprintf("^(%s) is not authenticated$", nameRegex), authenticationTest.isAuthenticated(false))
	sc.When(fmt.Sprintf("^(%s) submits a request to my service$", nameRegex), authenticationTest.submitsRequestForService)
	sc.When(fmt.Sprintf("^(%s) submits a wrong request to my service$", nameRegex), authenticationTest.submitsWrongRequest)
	sc.Then("^I should allow access to my service$", authenticationTest.shouldAllowAccess)
	sc.Then(fmt.Sprintf("^I should deny access saying (%s) with (%s) status code$", nameRegex, statusCodeRegex), authenticationTest.shouldDenyAccess)
}

func TestUserAuthentication(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "User authentication test suite",
		ScenarioInitializer: InitializeUserAuthenticationTest,
		Options: &godog.Options{
			Paths:    []string{"features"},
			Format:   "pretty",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run tests")
	}
}
