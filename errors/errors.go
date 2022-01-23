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
	ErrorFailToFetchItem       = errors.New("fail to fetch item") // temp
	ErrorFailToFetchRandomItem = errors.New("fail to fetch random item")
	ErrorFailToUploadItem      = errors.New("fail to upload item")
	ErrorFailToFindItem        = errors.New("fail to find item")
	ErrorInvalidItemIdentity   = errors.New("invalid item identity")

	// request errors
	ErrorFailToCreateRequest      = errors.New("fail to create request")
	ErrorRequestAlreadyExists     = errors.New("request already exists")
	ErrorRequestClosedItem        = errors.New("fail to request a closed item")
	ErrorFailToFindRequest        = errors.New("fail to find request")
	ErrorOnlyAcceptPendingRequest = errors.New("only accept a pending request")
	ErrorAcceptManyRequest        = errors.New("can only accept one request")
	ErrorDebug                    = errors.New("Debug")

	// message errors
	ErrorFailToFetchMessage = errors.New("fail to fetch message")
	ErrorFailToGetContacts  = errors.New("fail to get contacts")

	// review erros
	ErrorFailToAddReview  = errors.New("fail to add review")
	ErrorFailToGetReviews = errors.New("fail to get reviews")
)
