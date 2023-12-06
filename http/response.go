package http

type (
	// HTTPLambdaResponse HTTP lambda's response.
	HTTPLambdaResponse[R any] struct {
		Cookies         []string `json:"cookies,omitempty"`
		IsBase64Encoded bool     `json:"isBase64Encoded,omitempty"`
		StatusCode      int      `json:"statusCode,omitempty"`
		Headers         []string `json:"headers,omitempty"`
		Body            R        `json:"body,omitempty"`
	}
)
