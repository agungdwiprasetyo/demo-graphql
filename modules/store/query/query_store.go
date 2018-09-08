package query

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

type StoreQuery struct {
	Query
}

func NewStoreQuery(read *Query) *StoreQuery {
	store := new(StoreQuery)
	store.db = read.db
	return store
}

func (st *StoreQuery) GetByID(id int) (*model.Store, error) {
	store := new(model.Store)

	query := `SELECT id, name FROM stores WHERE id = $1`
	err := st.db.QueryRow(query, id).Scan(&store.ID, &store.Name)
	if err != nil {
		return nil, err
	}
	debug.Println(store)
	return store, nil
}
