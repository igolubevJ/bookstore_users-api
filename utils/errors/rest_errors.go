package errors

import "net/http"

// RESTErr - struct error for rest api
type RESTErr struct {
	Message string `json:"string"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError - return pointer for RESTErr with BadRequest Status and Error bad_request
func NewBadRequestError(message string) *RESTErr {
	return &RESTErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
