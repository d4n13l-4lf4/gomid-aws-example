package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/d4n13l-4lf4/gomid-aws-example/auth"
	"github.com/d4n13l-4lf4/gomid-aws-example/hello"
	"github.com/d4n13l-4lf4/gomid/middleware"
)

var allowedUsers = []string{"d4n13l-4lf4"}

func main() {
	greeter := hello.NewGreetingController(hello.Greet)
	chain := middleware.Wrap[func(context.Context, *events.APIGatewayProxyRequest) (string, error)](
		greeter,
	).
		Add(auth.AuthenticateUser(allowedUsers)).
		Build()

	lambda.Start(chain)
}
