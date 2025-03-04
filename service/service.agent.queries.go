package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s *Service) AddNewAgentAsCollector(ctx context.Context, param AddNewAgentAsCollectorParams) (Collector, error) {
	p := postgres.AddNewAgentAsCollectorParams{
		CollectorEmail:  param.CollectorEmail,
		FirstName:       param.FirstName,
		LastName:        param.LastName,
		DisplayName:     param.DisplayName,
		AvatarUrl:       param.AvatarUrl,
		LocationCity:    param.LocationCity,
		LocationState:   param.LocationState,
		LocationCountry: param.LocationCountry,
		LocationAddr:    param.LocationAddr,
	}

	tx, err := s.pgxPool.Begin(ctx)
	if err != nil {
		return Collector{}, fmt.Errorf("error begin-ing txn from pgxPool: error: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := s.dbQuery.WithTx(tx)
	collector, err := qtx.AddNewAgentAsCollector(ctx, p)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				if collector, err = s.dbQuery.GetCollector(
					context.Background(), param.CollectorEmail); err != nil {
					return Collector{}, err
				}
				return makeServiceCollector(collector)
			}
		}

		return Collector{}, fmt.Errorf("service.AddNewAgentAsCollector: error: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return Collector{}, fmt.Errorf("error commiting txn: error: %w", err)
	}

	return makeServiceCollector(collector)
}

func (s *Service) GetAllCollectors(ctx context.Context) ([]Collector, error) {

	q, err := s.dbQuery.GetAllCollectors(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]Collector, len(q))

	for i, v := range q {
		agentId, err := pgtypeUUIDToGoogleUUID(q[i].AgentID)
		if err != nil {
			return nil, fmt.Errorf("GetAllCollectors: err formating agentID for %v; error: %w", v.AgentID.String(), err)
		}
		res[i] = Collector{
			CollectorEmail:  v.CollectorEmail,
			FirstName:       v.FirstName,
			LastName:        v.LastName,
			DisplayName:     v.DisplayName,
			AvatarUrl:       v.AvatarUrl,
			LocationCity:    v.LocationCity,
			LocationState:   v.LocationState,
			LocationCountry: v.LocationCountry,
			LocationAddr:    v.LocationAddr,
			AgentID:         agentId,
			CollectorRole:   Collectorrole(v.CollectorRole),
		}
	}
	return res, nil
}
