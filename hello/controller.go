package hello

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/d4n13l-4lf4/gomid-aws-example/http"
)

type (
	GreetingRequest struct {
		Name string `json:"name,omitempty"`
	}

	Controller struct {
		Greeter func(context.Context, string) (string, error)
	}
)

func NewGreetingController(greeter Greeter) *Controller {
	return &Controller{
		Greeter: greeter,
	}
}

func (lg *Controller) Greet(
	ctx context.Context,
	event *events.APIGatewayProxyRequest) (*lambda.LambdaResponse[string], error) {
	var greetingRequest GreetingRequest

	err := json.Unmarshal([]byte(event.Body), &greetingRequest)
	if err != nil {
		//nolint:nilerr
		return &lambda.LambdaResponse[string]{
			StatusCode: http.StatusBadRequest,
			Body:       "bad request",
		}, nil
	}

	name := strings.TrimSpace(greetingRequest.Name)

	if strings.EqualFold(name, "") {
		return &lambda.LambdaResponse[string]{
			StatusCode: http.StatusBadRequest,
			Body:       "bad request",
		}, nil
	}

	greeting, err := lg.Greeter(ctx, name)
	if err != nil {
		// nolint:nilerr
		return &lambda.LambdaResponse[string]{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return &lambda.LambdaResponse[string]{
		StatusCode: http.StatusOK,
		Body:       greeting,
	}, nil
}
