package query

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/jmoiron/sqlx"
)

type Query struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

type (
	Store interface {
		FindAll() ([]model.Store, error)
		FindByID(int) (*model.Store, error)
	}

	Product interface {
		FindAll() ([]model.Product, error)
		FindByStoreID(int) ([]model.Product, error)
		FindByID(int) (*model.Product, error)
	}
)

func NewQuery(db *sqlx.DB) *Query {
	return &Query{db: db}
}
