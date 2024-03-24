package response_status

import "net/http"

type (
	StatusCode int
	StatusMsg  string
)

const (
	Ok                  StatusCode = http.StatusOK
	Created             StatusCode = http.StatusCreated
	BadRequest          StatusCode = http.StatusBadRequest
	InternalError       StatusCode = http.StatusInternalServerError
	UnprocessableEntity StatusCode = http.StatusUnprocessableEntity
	NotFound            StatusCode = http.StatusNotFound
)

const (
	SuccessMsg                StatusMsg = "success"
	InternalErrMsg            StatusMsg = "internal error"
	InvalidRequestErrMsg      StatusMsg = "invalid request"
	UnprocessableEntityErrMsg StatusMsg = "Inválid JSON"
	NotFoundErrMsg            StatusMsg = "Not Found"
)
