package models

type OrderItem struct {
	Dish     `json:"dish"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID         int    `json:"id"`
	Restaurant int    `json:"restaurant"`
	Items      string `json:"items"`
}

type CreateOrder struct {
	Restaurant int         `json:"restaurant"`
	Items      []OrderItem `json:"items"`
}
