package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/repository"
	"github.com/graphql-go/graphql"
)

type StoreUsecase interface {
	GetAllStore() *graphql.Field
	GetStoreByID() *graphql.Field
	GetAllProduct() *graphql.Field
	GetProductByID() *graphql.Field
	SaveStore(*model.Store) error
}

type storeUsecase struct {
	read              *query.Query
	storeQuery        query.Store
	productQuery      query.Product
	write             *repository.Repository
	storeRepository   repository.Store
	productRepository repository.Product
}

func NewStoreUsecase(read *query.Query, write *repository.Repository) StoreUsecase {
	return &storeUsecase{
		read:              read,
		storeQuery:        query.NewStoreQuery(read),
		productQuery:      query.NewProductQuery(read),
		write:             write,
		storeRepository:   repository.NewStoreRepository(write),
		productRepository: repository.NewProductRepository(write),
	}
}
