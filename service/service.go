package service

import (
	"beltech/service/internal/repository/postgres"
	"context"

	"github.com/jackc/pgx/v5"
)

//go:generate sqlc generate

type Service struct {
	dbQuery *postgres.Queries
	pgxPool PostgresDb
}

type PostgresDb interface {
	postgres.DBTX
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewService(postgresDb PostgresDb) *Service {
	var s Service
	s.dbQuery = postgres.New(postgresDb)
	s.pgxPool = postgresDb
	return &s
}
