package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// import (
// 	"beltech/service/internal/repository/postgres"
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v5/pgtype"
// )

// func (s *Service) AddNewTempWasteFromCollector(ctx context.Context, param AddNewWasteFromCollectorParams) (uuid.UUID, error) {

// 	pgUUID, err := s.dbQuery.AddNewWasteFromCollector(ctx, makeAddNewWasteFromCollectorParams(param))
// 	if err != nil {
// 		return uuid.Nil, err
// 	}
// 	return pgtypeUUIDToGoogleUUID(pgUUID)
// }

// func (s *Service) GetWasteFromCollectorByAgentIDAndDate(ctx context.Context, agentID uuid.UUID, date time.Time) ([]GetWasteFromCollectorByAgentIDAndDateRow, error) {
// 	pgUUID, err := googleUUIDToPgtypeUUID(agentID)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing agentID: %w", err)
// 	}

// 	res, err := s.dbQuery.GetWasteFromCollectorByAgentIDAndDate(ctx, postgres.GetWasteFromCollectorByAgentIDAndDateParams{
// 		AgentID: pgUUID,
// 		Column2: pgtype.Date{Time: date, Valid: true},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return makeGetWasteFromCollectorByAgentIDAndDateRow(res)
// }

func (s *Service) AddNewAgentAsCollector(ctx context.Context, param AddNewAgentAsCollectorParams) (uuid.UUID, error) {
	tx, err := s.pgxPool.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	qtx := s.dbQuery.WithTx(tx)

	p := postgres.AddNewAgentAsCollectorParams{
		AgentType:      postgres.Identifyas(string(param.AgentType)),
		City:           param.City,
		LocationState:  param.LocationState,
		StMakepoint:    param.StMakepoint,
		StMakepoint_2:  param.StMakepoint_2,
		CollectorEmail: param.CollectorEmail,
		CollectorRole:  postgres.Collectorrole(string(param.CollectorRole)),
		FirstName:      param.FirstName,
		LastName:       param.LastName,
		DisplayName:    param.DisplayName,
		AvatarUrl:      param.AvatarUrl,
	}

	pgUUID, err := qtx.AddNewAgentAsCollector(ctx, p)
	if err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	pgTypeUUID, ok := pgUUID.(pgtype.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("unexpected type for UUID: %T", pgUUID)
	}

	return pgtypeUUIDToGoogleUUID(pgTypeUUID)
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
			return nil, fmt.Errorf("err formating agentID for %v; error: %w", v.AgentID.String(), err)
		}
		res[i] = Collector{
			CollectorEmail: v.CollectorEmail,
			FirstName:      v.FirstName,
			LastName:       v.LastName,
			DisplayName:    v.DisplayName,
			AvatarUrl:      v.AvatarUrl,
			CollectorRole:  Collectorrole(string(q[i].CollectorRole)),
			AgentID:        agentId,
		}
	}

	return res, nil
}
