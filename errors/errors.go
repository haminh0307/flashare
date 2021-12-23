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

	// item errors
	ErrorFailToFetchItem  = errors.New("fail to fetch item") // temp
	ErrorFailToUploadItem = errors.New("fail to upload item")

	// message errors
	ErrorFailToFetchMessage = errors.New("fail to fetch message")
)
