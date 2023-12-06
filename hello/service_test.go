package hello_test

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/d4n13l-4lf4/gomid-aws-example/hello"
	"github.com/d4n13l-4lf4/gomid-aws-example/internal"
	"github.com/stretchr/testify/assert"
)

const (
	greetingTestKey = internal.TestKey("greetingTestKey")
)

type (
	GreetingTest struct {
		greeter hello.Greeter
		name    string
	}
)

func (gt *GreetingTest) wantToGreetSomeone(ctx context.Context) (context.Context, error) {
	test := ctx.Value(greetingTestKey).(*GreetingTest)
	test.greeter = hello.Greet

	return context.WithValue(ctx, greetingTestKey, test), nil
}

func (gt *GreetingTest) receiveGreetingRequest(ctx context.Context, name string) (context.Context, error) {
	test := ctx.Value(greetingTestKey).(*GreetingTest)
	test.name = name

	return context.WithValue(ctx, greetingTestKey, test), nil
}

func (gt *GreetingTest) shouldGreetSaying(ctx context.Context, expectedGreeting string) error {
	asserter := internal.NewAsserter()
	assertions := assert.New(asserter)

	test := ctx.Value(greetingTestKey).(*GreetingTest)

	greeting, err := test.greeter(ctx, test.name)

	assertions.Equal(expectedGreeting, greeting)
	assertions.NoError(err)

	return asserter.Error()
}

func InitializeServiceTests(ctx *godog.ScenarioContext) {
	var greetingTest GreetingTest

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		greetingTest := GreetingTest{}

		return context.WithValue(ctx, greetingTestKey, &greetingTest), nil
	})

	ctx.Given("^I want to greet ([A-Za-z0-9\\s]+)$", greetingTest.wantToGreetSomeone)
	ctx.When("^I receive a greeting request for ([A-Za-z0-9\\s]+)$", greetingTest.receiveGreetingRequest)
	ctx.Then("^I should greet saying ([A-Za-z0-9!\\s]+)$", greetingTest.shouldGreetSaying)

}

func TestGreetingService(t *testing.T) {
	suite := godog.TestSuite{
		Name:                "GreetingService test suite",
		ScenarioInitializer: InitializeServiceTests,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features/service"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run tests")
	}
}
