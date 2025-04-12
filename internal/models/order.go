package models

type Order struct {
	Dish     `json:"dish"`
	Quantity int `json:"quantity"`
}
