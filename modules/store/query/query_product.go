package query

import "github.com/agungdwiprasetyo/demo-graphql/modules/store/model"

type ProductQuery struct {
	// db *sqlx.DB
}

func NewProductQuery() *ProductQuery {
	return new(ProductQuery)
}

func (st *ProductQuery) GetByStoreID(storeID int) ([]model.Product, error) {
	products := []model.Product{
		{
			ID: 1, Name: "Hape seken",
		},
		{
			ID: 2, Name: "Hape baru",
		},
	}

	return products, nil
}
