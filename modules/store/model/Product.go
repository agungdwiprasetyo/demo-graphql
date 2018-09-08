package model

type Product struct {
	ID      int    `json:"product_id"`
	StoreID int    `json:"store_id,omitempty"`
	Name    string `json:"product_name"`
}
