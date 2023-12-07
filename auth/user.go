package auth

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/d4n13l-4lf4/gomid-aws-example/http"
	"github.com/d4n13l-4lf4/gomid/middleware"
	"github.com/samber/lo"
)

const (
	HeaderUsername = "X-D4N-USERNAME"
)

func AuthenticateUser(users []string) middleware.Middleware {
	return func(nxt middleware.Next) middleware.Next {
		return func(ctx context.Context, rawEvent any) (any, error) {
			event, ok := rawEvent.(*events.APIGatewayProxyRequest)
			if !ok {
				return &lambda.LambdaResponse[string]{
					Body:       "wrong request for authorization",
					StatusCode: http.StatusForbidden,
				}, nil
			}

			key, _ := lo.FindKeyBy[string](event.Headers, func(key, value string) bool {
				return strings.EqualFold(key, HeaderUsername)
			})

			username := event.Headers[key]

			isKnownUser := slices.Contains(users, username)

			if !isKnownUser {
				return &lambda.LambdaResponse[string]{
					Body:       fmt.Errorf("%s is not allowed to access this resource", username).Error(),
					StatusCode: http.StatusForbidden,
				}, nil
			}

			return nxt(ctx, rawEvent)
		}
	}
}
