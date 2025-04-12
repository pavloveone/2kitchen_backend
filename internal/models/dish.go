package models

type Dish struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Protein     float64 `json:"protein"`
	Fat         float64 `json:"fat"`
	Carbs       float64 `json:"carbs"`
	Calories    float64 `json:"calories"`
	Restaurant  int     `json:"restaurant"`
}

type ModificationDish struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Protein     float64 `json:"protein"`
	Fat         float64 `json:"fat"`
	Carbs       float64 `json:"carbs"`
	Calories    float64 `json:"calories"`
	Restaurant  int     `json:"restaurant"`
}
