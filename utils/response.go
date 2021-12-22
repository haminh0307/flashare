package utils

import ()

type DataResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
