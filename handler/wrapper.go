package handler

type Wrapper struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
	StackTrace string      `json:"stackTrace,omitempty"`
}
