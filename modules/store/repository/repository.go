package repository

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

type (
	Store interface {
		Save(*model.Store) <-chan error
	}

	Product interface {
		Save(*model.Product) <-chan error
	}
)

var tx *sqlx.Tx

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) StartTransaction() {
	tx, _ = repo.db.Beginx()
}

func (repo *Repository) Rollback() {
	if tx != nil {
		tx.Rollback()
	}
	tx = nil
}

func (repo *Repository) Commit() {
	if tx != nil {
		tx.Commit()
	}
	tx = nil
}
