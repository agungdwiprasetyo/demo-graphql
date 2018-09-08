package model

type Store struct {
	ID       int       `json:"store_id"`
	Name     string    `json:"store_name"`
	Products []Product `json:"products,omitempty"`
}
