package query

import "github.com/agungdwiprasetyo/demo-graphql/modules/store/model"

type ProductQuery struct {
	Query
}

func NewProductQuery(read *Query) *ProductQuery {
	product := new(ProductQuery)
	product.db = read.db
	return product
}

func (st *ProductQuery) GetByStoreID(storeID int) ([]model.Product, error) {
	query := `SELECT id, name FROM products WHERE store_id = $1`
	rows, err := st.db.Query(query, storeID)
	if err != nil {
		return nil, err
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
