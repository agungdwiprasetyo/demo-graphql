package query

import "github.com/agungdwiprasetyo/demo-graphql/modules/store/model"

type StoreQuery struct {
	// db *sqlx.DB
}

func NewStoreQuery() *StoreQuery {
	return new(StoreQuery)
}

func (st *StoreQuery) GetByID(id int) (*model.Store, error) {
	store := new(model.Store)
	store.ID = id
	store.Name = "Pantau Cocok Bayar"
	return store, nil
}
