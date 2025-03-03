package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) AddNewAgentAsCollector(ctx context.Context, param AddNewAgentAsCollectorParams) (uuid.UUID, error) {
	p := postgres.AddNewAgentAsCollectorParams{
		CollectorEmail:  pgtype.Text{String: param.CollectorEmail, Valid: true},
		FirstName:       param.FirstName,
		LastName:        param.LastName,
		DisplayName:     param.DisplayName,
		AvatarUrl:       param.AvatarUrl,
		LocationCity:    param.LocationCity,
		LocationState:   param.LocationState,
		LocationCountry: param.LocationCountry,
		StMakepoint:     param.StMakepoint,
		StMakepoint_2:   param.StMakepoint_2,
		CollectorRole:   postgres.Collectorrole(string(param.CollectorRole)),
	}

	pgUUID, err := s.dbQuery.AddNewAgentAsCollector(ctx, p)
	if err != nil {
		return uuid.Nil, err
	}

	return pgtypeUUIDToGoogleUUID(pgUUID)
}

func (s *Service) GetAllCollectors(ctx context.Context) ([]AgentCollector, error) {

	q, err := s.dbQuery.GetAllCollectors(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]AgentCollector, len(q))

	for i, v := range q {
		agentId, err := pgtypeUUIDToGoogleUUID(q[i].AgentID)
		if err != nil {
			return nil, fmt.Errorf("err formating agentID for %v; error: %w", v.AgentID.String(), err)
		}
		res[i] = AgentCollector{
			AgentID:         agentId,
			CollectorEmail:  pgtextToString(v.CollectorEmail),
			FirstName:       v.FirstName,
			LastName:        v.LastName,
			DisplayName:     v.DisplayName,
			AvatarUrl:       v.AvatarUrl,
			LocationCity:    v.LocationCity,
			LocationState:   v.LocationState,
			LocationCountry: v.LocationCountry,
			LocationAddr:    v.LocationAddr,
			CollectorRole:   Collectorrole(string(v.CollectorRole)),
			CreatedAt:       v.CreatedAt.Time,
			IsDeleted:       v.IsDeleted.Bool,
			DeletedAt:       v.DeletedAt.Time,
		}
	}
	return res, nil
}
