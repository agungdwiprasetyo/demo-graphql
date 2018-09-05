package query

import "github.com/agungdwiprasetyo/demo-graphql/modules/store/model"

type (
	Store interface {
		GetByID(int) (*model.Store, error)
	}

	Product interface {
		GetByStoreID(int) ([]model.Product, error)
	}
)
