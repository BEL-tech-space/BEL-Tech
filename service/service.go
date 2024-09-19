package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
)

//go:generate sqlc generate


type Service struct {
	dbQuery *postgres.Queries
}

func NewService() *Service {
	return &Service{}
}

type User struct {
	postgres.User
}

func (s *Service) CreateOrUpdateUser(ctx context.Context, user User) error {
	
	var err error
    if err = s.dbQuery.UpsertUser(ctx, convertToUpsertUserParams(user)); err == nil {
		return nil
	}

	// handle error
	return err
}

