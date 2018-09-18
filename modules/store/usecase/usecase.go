package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/graphql-go/graphql"
)

// StoreUsecase abstraction
type StoreUsecase interface {
	GetAllStore() *graphql.Field
	GetStoreByID() *graphql.Field
	GetAllProduct() *graphql.Field
	GetProductByID() *graphql.Field
	SaveStore(stores *model.Store) error
}
