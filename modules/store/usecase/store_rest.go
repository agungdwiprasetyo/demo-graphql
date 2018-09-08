package usecase

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
)

func (uc *storeUsecase) SaveStore(data *model.Store) error {
	uc.write.StartTransaction()

	err := <-uc.storeRepository.Save(data)
	if err != nil {
		uc.write.Rollback()
		return err
	}

	for _, product := range data.Products {
		tmp := product
		tmp.StoreID = data.ID
		err := <-uc.productRepository.Save(&tmp)
		if err != nil {
			uc.write.Rollback()
			return err
		}
	}

	uc.write.Commit()
	return nil
}
