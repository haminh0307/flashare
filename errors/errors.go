package flashare_errors

import (
	"errors"
)

var (
	// server errors
	ErrorInternalServerError = errors.New("internal server error")
	ErrorInvalidParameters   = errors.New("invalid parameters")

	// auth errors
	ErrorInvalidCredentials = errors.New("invalid credentials provided")
	ErrorEmailAlreadyExists = errors.New("email already exists")
)
