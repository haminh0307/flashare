package utils

import ()

type DataResponse struct {
	Status string      `json:"status"` // okay, fail
	Data   interface{} `json:"data"`
}
