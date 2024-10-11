package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func (s *Service) CreateOrUpdateUser(ctx context.Context, user User) (User, error) {

	u, err := s.dbQuery.UpsertUser(ctx, makeUpsertUserParams(user))
	if err == nil {
		return broadendUser(u, user), nil
	}

	// handle error
	return User{}, err
}

func (s *Service) AddNewTempWasteFromCollector(ctx context.Context, param AddNewTempWasteFromCollectorParams) (uuid.UUID, error) {

	pgUUID, err := s.dbQuery.AddNewTempWasteFromCollector(ctx, makeAddNewTempWasteFromCollectorParams(param))
	if err != nil {
		return uuid.Nil, err
	}
	return pgtypeUUIDToGoogleUUID(pgUUID)
}

func (s *Service) GetWasteFromCollectorByAgentIDAndDate(ctx context.Context, agentID uuid.UUID, date time.Time) ([]GetWasteFromCollectorByAgentIDAndDateRow, error) {
	pgUUID, err := googleUUIDToPgtypeUUID(agentID)
	if err != nil {
		return nil, fmt.Errorf("error parsing agentID: %w", err)
	}

	res, err := s.dbQuery.GetWasteFromCollectorByAgentIDAndDate(ctx, postgres.GetWasteFromCollectorByAgentIDAndDateParams{
		AgentID: pgUUID,
		Column2: pgtype.Date{Time: date, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return makeGetWasteFromCollectorByAgentIDAndDateRow(res)
}

func (s *Service) AddNewAgentAndCollector(ctx context.Context, param AddNewAgentParams) (uuid.UUID, error) {
	tx, err := s.pgxPool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	qtx := s.dbQuery.WithTx(tx)
	p, err := MakeAddNewAgentParams(param)
	if err != nil {
		return uuid.Nil, fmt.Errorf("some params were wrong: %w", err)
	}
	pgUUID, err := qtx.AddNewAgentAndCollector(ctx, p)
	if err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	return pgtypeUUIDToGoogleUUID(pgUUID)
}

func (s *Service) GetAllCollectorsEmailAndAgentID(ctx context.Context) ([]GetAllCollectorsEmailAndAgentIDRow, error) {

	q, err := s.dbQuery.GetAllCollectorsEmailAndAgentID(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]GetAllCollectorsEmailAndAgentIDRow, 0, len(q))

	for i := 0; i < len(q); i++ {
		res[i].Email = q[i].Email
		res[i].Role = q[i].Role

		res[i].AgentID, err = pgtypeUUIDToGoogleUUID(q[i].AgentID)
		if err != nil {
			return nil, fmt.Errorf("error parsing pgUUID to ggUUID: %w", err)
		}
	}

	return res, nil
}
