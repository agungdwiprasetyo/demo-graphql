package model

type Store struct {
	ID   int    `json:"store_id"`
	Name string `json:"store_name"`
}

type Product struct {
	ID   int    `json:"product_id"`
	Name string `json:"product_name"`
}
