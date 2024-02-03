package constants

import "net/http"

const (
	BadRequestMessage   = "Invalid request, please ensure that all the required parameters are passed"
	InternalServerError = "Oops! Something went wrong, please try again."
	Unauthorised        = "Oops! Seems like you do not have access for this action."
)

var StatusMessageMap = map[int]string{
	http.StatusBadRequest:          BadRequestMessage,
	http.StatusInternalServerError: InternalServerError,
	http.StatusUnauthorized:        Unauthorised,
}
