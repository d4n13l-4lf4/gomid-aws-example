package hello

import (
	"context"
	"net/http"
	"strings"

	lambda "github.com/d4n13l-4lf4/gomid-aws-example/http"
)

type (
	GreetingRequest struct {
		Name string `json:"name,omitempty"`
	}

	GreetingController struct {
		Greeter func(context.Context, string) (string, error)
	}
)

func NewGreetingController(greeter Greeter) *GreetingController {
	return &GreetingController{
		Greeter: greeter,
	}
}

func (lg *GreetingController) Greet(ctx context.Context, request *GreetingRequest) (*lambda.HTTPLambdaResponse[string], error) {
	name := strings.TrimSpace(request.Name)

	if strings.EqualFold(name, "") {
		return &lambda.HTTPLambdaResponse[string]{
			StatusCode: http.StatusBadRequest,
			Body:       "bad request",
		}, nil
	}

	greeting, err := lg.Greeter(ctx, request.Name)
	if err != nil {
		return &lambda.HTTPLambdaResponse[string]{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return &lambda.HTTPLambdaResponse[string]{
		StatusCode: http.StatusOK,
		Body:       greeting,
	}, nil
}
