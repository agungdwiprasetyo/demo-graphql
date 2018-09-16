package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	repositoryDecorator "github.com/agungdwiprasetyo/demo-graphql/modules/store/repository"
	"github.com/graphql-go/graphql"
)

// StoreUsecase abstraction
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
	repository        *repositoryDecorator.Repository
	storeRepository   repositoryDecorator.Store
	productRepository repositoryDecorator.Product
}

func NewStoreUsecase(read *query.Query, repository *repositoryDecorator.Repository) StoreUsecase {
	return &storeUsecase{
		read:              read,
		storeQuery:        query.NewStoreQuery(read),
		productQuery:      query.NewProductQuery(read),
		repository:        repository,
		storeRepository:   repositoryDecorator.NewStoreRepository(repository),
		productRepository: repositoryDecorator.NewProductRepository(repository),
	}
}
