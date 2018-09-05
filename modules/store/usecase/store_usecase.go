package usecase

import "github.com/graphql-go/graphql"

type StoreUsecase interface {
	GetStore() (graphql.Schema, error)
}
