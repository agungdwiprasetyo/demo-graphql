package repository

import (
	"github.com/agungdwiprasetyo/demo-graphql/modules/store/model"
	"github.com/jmoiron/sqlx"
)

// Abstraction method

type (
	// Store abstraction
	Store interface {
		Save(*model.Store) <-chan error
	}

	// Product abstraction
	Product interface {
		Save(*model.Product) <-chan error
	}
)

// Repository parent domain
type Repository struct {
	db *sqlx.DB
}

var tx *sqlx.Tx

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) StartTransaction() {
	t, err := repo.db.Beginx()
	if err != nil {
		panic(err)
	}
	tx = t
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
