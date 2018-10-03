package repository

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/agungdwiprasetyo/go-utils/debug"
	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	Repository
}

func NewProductRepository(write *Repository) Product {
	store := new(productRepository)
	store.db = write.db
	return store
}

func (repo *productRepository) Save(data *model.Product) <-chan error {
	output := make(chan error)

	go func() {
		defer close(output)

		var (
			stmt *sqlx.Stmt
			err  error
		)

		query := `INSERT INTO products (id, store_id, name) VALUES ($1, $2, $3)`
		if tx != nil {
			debug.Println("Using transaction")
			stmt, err = tx.Preparex(query)
		} else {
			stmt, err = repo.db.Preparex(query)
		}

		if err != nil {
			output <- err
			return
		}
		_, err = stmt.Exec(data.ID, data.StoreID, data.Name)
		if err != nil {
			output <- err
			return
		}
		output <- nil
	}()

	return output
}
