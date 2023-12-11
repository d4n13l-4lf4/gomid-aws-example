package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/d4n13l-4lf4/gomid-aws-example/auth"
	"github.com/d4n13l-4lf4/gomid-aws-example/hello"
	httpLambda "github.com/d4n13l-4lf4/gomid-aws-example/http"
	"github.com/d4n13l-4lf4/gomid/middleware"
)

// nolint: gochecknoglobals
var allowedUsers = []string{"d4n13l-4lf4"}

type (
	// Greet greeting controller type.
	Greet func(context.Context, *events.APIGatewayProxyRequest) (*httpLambda.LambdaResponse[string], error)
)

func main() {
	greeter := hello.NewGreetingController(hello.Greet)
	chain := middleware.Wrap[Greet](
		greeter.Greet,
	).
		Add(auth.AuthenticateUser(allowedUsers)).
		Build()

	lambda.Start(chain)
}
