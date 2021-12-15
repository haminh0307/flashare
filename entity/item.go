package entity

import (

)

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
	From string `json:"from"`
	To string `json:"to"`
}