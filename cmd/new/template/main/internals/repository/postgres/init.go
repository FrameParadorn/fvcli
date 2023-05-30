package postgres

import "git.robodev.co/newfinvest/platform/common-service/internals/infrastructure/database"

type PostgresRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Interface {
	return &PostgresRepository{db: db}
}
