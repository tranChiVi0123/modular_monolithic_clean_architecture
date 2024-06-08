package errors_handler

import "fmt"

type errorReponse struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (e *errorReponse) Error() string {
	return fmt.Sprintf("Code: %d, Title: %s, Message: %s", e.Code, e.Title, e.Message)
}

func New400ErrorResponse(err error) *errorReponse {
	return &errorReponse{
		Code:    400,
		Title:   "Bad Request",
		Message: err.Error(),
	}
}

func New404ErrorResponse() *errorReponse {
	return &errorReponse{
		Code:    404,
		Title:   "Not Found",
		Message: "Resource not found",
	}
}

func New500ErrorResponse(err error) *errorReponse {
	return &errorReponse{
		Code:    500,
		Title:   "Internal Server Error",
		Message: err.Error(),
	}
}

func New401ErrorResponse() *errorReponse {
	return &errorReponse{
		Code:    401,
		Title:   "Unauthorized",
		Message: "Unauthorized access",
	}
}

func New403ErrorResponse() *errorReponse {
	return &errorReponse{
		Code:    403,
		Title:   "Forbidden",
		Message: "Forbidden access",
	}
}
