package models

type MenuItem struct {
	ID       string
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type Menu struct {
	Items []MenuItem
}
