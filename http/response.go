package http

type (
	// LambdaResponse HTTP lambda's response.
	LambdaResponse[R any] struct {
		Cookies         []string `json:"cookies,omitempty"`
		IsBase64Encoded bool     `json:"isBase64Encoded,omitempty"`
		StatusCode      int      `json:"statusCode,omitempty"`
		Headers         []string `json:"headers,omitempty"`
		Body            R        `json:"body,omitempty"`
	}
)
