package flashare_errors

import (
	"errors"
)

var (
	// server errors
	ErrorInternalServerError = errors.New("internal server error")
	ErrorInvalidParameters   = errors.New("invalid parameters")

	// mongo object ID errors
	ErrorInvalidObjectID = errors.New("invalid object ID")

	// auth errors
	ErrorInvalidCredentials = errors.New("invalid credentials provided")
	ErrorEmailAlreadyExists = errors.New("email already exists")

	// profile errors
	ErrorUserNotExists = errors.New("user not exists")

	// item errors
	ErrorFailToFetchItem     = errors.New("fail to fetch item") // temp
	ErrorFailToUploadItem    = errors.New("fail to upload item")
	ErrorFailToFindItem      = errors.New("fail to find item")
	ErrorInvalidItemIdentity = errors.New("invalid item identity")

	// request errors
	ErrorFailToCreateRequest  = errors.New("fail to create request")
	ErrorRequestAlreadyExists = errors.New("request already exists")
	ErrorDebug                = errors.New("Debug")

	// message errors
	ErrorFailToFetchMessage = errors.New("fail to fetch message")
)
