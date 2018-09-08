package repository

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/jmoiron/sqlx"
)

type StoreRepository struct {
	Repository
}

func NewStoreRepository(write *Repository) *StoreRepository {
	store := new(StoreRepository)
	store.db = write.db
	store.tx = write.tx
	return store
}

func (repo *StoreRepository) Save(data *model.Store) <-chan error {
	output := make(chan error)

	go func() {
		defer close(output)

		var (
			stmt *sqlx.Stmt
			err  error
		)
		query := `INSERT INTO stores (id, name) VALUES ($1, $2)`
		if tx != nil {
			stmt, err = tx.Preparex(query)
		} else {
			stmt, err = repo.db.Preparex(query)
		}

		if err != nil {
			output <- err
			return
		}

		_, err = stmt.Exec(data.ID, data.Name)
		if err != nil {
			output <- err
			return
		}
		output <- nil

	}()

	return output
}
