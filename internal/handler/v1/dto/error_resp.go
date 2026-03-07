package dto

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func NewErrorResponse(errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		ErrorMessage: errorMessage,
	}
}
