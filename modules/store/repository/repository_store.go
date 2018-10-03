package repository

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/jmoiron/sqlx"
)

type storeRepository struct {
	Repository
}

func NewStoreRepository(write *Repository) Store {
	store := new(storeRepository)
	store.db = write.db
	return store
}

func (repo *storeRepository) Save(data *model.Store) <-chan error {
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
