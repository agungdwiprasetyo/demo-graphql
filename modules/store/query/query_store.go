package query

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
)

type storeQuery struct {
	Query
}

func NewStoreQuery(read *Query) Store {
	store := new(storeQuery)
	store.db = read.db
	return store
}

func (st *storeQuery) FindAll() ([]model.Store, error) {
	stores := make([]model.Store, 0)
	query := `SELECT id, name FROM stores`
	rows, err := st.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var store model.Store
		if err := rows.Scan(&store.ID, &store.Name); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	return stores, nil
}

func (st *storeQuery) FindByID(id int) (*model.Store, error) {
	store := new(model.Store)
	query := `SELECT id, name FROM stores WHERE id = $1`
	err := st.db.QueryRow(query, id).Scan(&store.ID, &store.Name)
	if err != nil {
		return nil, err
	}
	return store, nil
}
