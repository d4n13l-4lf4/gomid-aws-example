package hello_test

import (
	"context"
	"errors"
	"testing"

	"github.com/cucumber/godog"
	"github.com/d4n13l-4lf4/gomid-aws-example/hello"
	lambda "github.com/d4n13l-4lf4/gomid-aws-example/http"
	"github.com/d4n13l-4lf4/gomid-aws-example/internal"
	mockHello "github.com/d4n13l-4lf4/gomid-aws-example/mocks/hello"
	"github.com/stretchr/testify/assert"
)

const (
	controllerTestKey = internal.TestKey("greeting controller test")
	emptyString       = ""
)

type (
	GreetingControllerTest struct {
		greeter          *mockHello.Greeter
		controller       *hello.GreetingController
		greetingRequest  *hello.GreetingRequest
		asserter         *internal.Asserter
		greetingResponse *lambda.HTTPLambdaResponse[string]
		err              error
	}
)

func (gt *GreetingControllerTest) receiveGreetingRequest(ctx context.Context, name string) (context.Context, error) {
	gt.asserter = internal.NewAsserter()
	gt.greetingRequest = &hello.GreetingRequest{
		Name: name,
	}
	gt.greeter = &mockHello.Greeter{}
	gt.controller = hello.NewGreetingController(gt.greeter.Execute)

	return ctx, nil
}

func (gt *GreetingControllerTest) receiveInvalidGreetingRequest(ctx context.Context) (context.Context, error) {
	return gt.receiveGreetingRequest(ctx, emptyString)
}

func (gt *GreetingControllerTest) greetHello(ctx context.Context, greeting string) (context.Context, error) {
	gt.greeter.EXPECT().
		Execute(ctx, gt.greetingRequest.Name).
		Return(greeting, nil).
		Maybe()

	gt.greetingResponse, gt.err = gt.controller.Greet(ctx, gt.greetingRequest)

	return ctx, nil
}

func (gt *GreetingControllerTest) shouldGreetSuccessfully(_ context.Context, greeting string, statusCode int) error {
	assertions := assert.New(gt.asserter)

	expectedResponse := &lambda.HTTPLambdaResponse[string]{
		StatusCode: statusCode,
		Body:       greeting,
	}

	assertions.Equal(gt.greetingResponse, expectedResponse)
	assertions.NoError(gt.err)

	gt.greeter.AssertExpectations(gt.asserter)

	return gt.asserter.Error()
}

func (gt *GreetingControllerTest) failToGreet(ctx context.Context, errMsg string) error {
	gt.greeter.EXPECT().
		Execute(ctx, gt.greetingRequest.Name).
		Return(emptyString, errors.New(errMsg)).
		Once()

	gt.greetingResponse, gt.err = gt.controller.Greet(ctx, gt.greetingRequest)

	return nil
}

func (gt *GreetingControllerTest) shouldGetAnError(_ context.Context, errMsg string, statusCode int) error {
	assertions := assert.New(gt.asserter)
	expectedResponse := &lambda.HTTPLambdaResponse[string]{
		StatusCode: statusCode,
		Body:       errMsg,
	}

	assertions.Equal(gt.greetingResponse, expectedResponse)
	assertions.NoError(gt.err)
	gt.greeter.AssertExpectations(gt.asserter)

	return gt.asserter.Error()
}

func InitializeControllerTests(ctx *godog.ScenarioContext) {
	var greetingTest GreetingControllerTest
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		greetingTest = GreetingControllerTest{}

		return context.WithValue(ctx, greetingTestKey, &greetingTest), nil
	})

	ctx.Given("^I receive a greeting request for ([A-Za-z0-9\\s]+)$", greetingTest.receiveGreetingRequest)
	ctx.Given("^I receive an invalid greeting request$", greetingTest.receiveInvalidGreetingRequest)

	ctx.When("^I greet as ([A-Za-z0-9!\\s]*)$", greetingTest.greetHello)
	ctx.Then("^I should greet ([A-Za-z0-9\\s]+) successfully with (\\d{3})$", greetingTest.shouldGreetSuccessfully)

	ctx.When("^Greeting fails with error ([A-Za-z0-9\\s]+)$", greetingTest.failToGreet)
	ctx.Then("^I should get an error ([A-Za-z0-9\\s]+) with (\\d{3})$", greetingTest.shouldGetAnError)

}

func TestHelloController(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "GreetingController test suite",
		ScenarioInitializer: InitializeControllerTests,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/controller"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run tests")
	}
}
