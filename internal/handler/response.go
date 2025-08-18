package handler

type errorRepsonse struct {
	ErrorMessage string `json:"error_message"`
}

func NewErrorResponse(errorMessage string) *errorRepsonse {
	return &errorRepsonse{
		ErrorMessage: errorMessage,
	}
}
