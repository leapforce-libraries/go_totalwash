package totalwash

// ErrorResponse stores general Ridder API error response
//
type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
