package query

import "github.com/agungdwiprasetyo/demo-graphql/modules/store/model"

type productQuery struct {
	Query
}

func NewProductQuery(read *Query) Product {
	product := new(productQuery)
	product.db = read.db
	return product
}

func (p *productQuery) FindAll() ([]model.Product, error) {
	query := `SELECT id, store_id, name FROM products`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	products := make([]model.Product, 0)
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.StoreID, &product.Name); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *productQuery) FindByStoreID(storeID int) ([]model.Product, error) {
	query := `SELECT id, name FROM products WHERE store_id = $1`
	rows, err := p.db.Query(query, storeID)
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

func (p *productQuery) FindByID(productID int) (*model.Product, error) {
	var product model.Product
	query := `SELECT id, store_id, name FROM products WHERE id = $1`
	err := p.db.QueryRow(query, productID).Scan(&product.ID, &product.StoreID, &product.Name)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
