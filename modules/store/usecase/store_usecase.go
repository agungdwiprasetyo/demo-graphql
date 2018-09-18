package usecase

import (
	queryDecorator "github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	repositoryDecorator "github.com/agungdwiprasetyo/demo-graphql/modules/store/repository"
)

type storeUsecase struct {
	query             *queryDecorator.Query
	storeQuery        queryDecorator.Store
	productQuery      queryDecorator.Product
	repository        *repositoryDecorator.Repository
	storeRepository   repositoryDecorator.Store
	productRepository repositoryDecorator.Product
}

func NewStoreUsecase(query *queryDecorator.Query, repository *repositoryDecorator.Repository) StoreUsecase {
	return &storeUsecase{
		query:             query,
		storeQuery:        queryDecorator.NewStoreQuery(query),
		productQuery:      queryDecorator.NewProductQuery(query),
		repository:        repository,
		storeRepository:   repositoryDecorator.NewStoreRepository(repository),
		productRepository: repositoryDecorator.NewProductRepository(repository),
	}
}
