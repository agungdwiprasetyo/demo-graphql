package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/query"
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/repository"
	"github.com/graphql-go/graphql"
)

type StoreUsecase interface {
	GetStore() (graphql.Schema, error)
	SaveStore(*model.Store) error
}

type storeUsecase struct {
	read              *query.Query
	storeQuery        query.Store
	productQuery      query.Product
	repo              *repository.Repository
	storeRepository   repository.Store
	productRepository repository.Product
}

func NewStoreUsecase(read *query.Query, repo *repository.Repository) StoreUsecase {
	return &storeUsecase{
		read:              read,
		storeQuery:        query.NewStoreQuery(read),
		productQuery:      query.NewProductQuery(read),
		repo:              repo,
		storeRepository:   repository.NewStoreRepository(repo),
		productRepository: repository.NewProductRepository(repo),
	}
}

func (uc *storeUsecase) SaveStore(data *model.Store) error {
	uc.repo.StartTransaction()

	err := <-uc.storeRepository.Save(data)
	if err != nil {
		uc.repo.Rollback()
		return err
	}

	for _, product := range data.Products {
		tmp := product
		tmp.StoreID = data.ID
		err := <-uc.productRepository.Save(&tmp)
		if err != nil {
			uc.repo.Rollback()
			return err
		}
	}

	uc.repo.Commit()
	return nil
}
