package configs

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cause   string `json:"cause,omitempty"`
}

func Success(status int, message string, data interface{}) *ResponseSuccess {
	return &ResponseSuccess{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
func Failed(status int, message string, cause string) *ResponseError {
	return &ResponseError{
		Status:  status,
		Message: message,
		Cause:   cause,
	}
}
