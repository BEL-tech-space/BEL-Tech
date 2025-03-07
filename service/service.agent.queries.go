package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
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
func (s *Service) AddNewAgentAsSuperCollector(ctx context.Context, param AddNewAgentAsSuperCollectorParams) (Collector, error) {
	p := postgres.AddNewAgentAsSuperCollectorParams{
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
	collector, err := qtx.AddNewAgentAsSuperCollector(ctx, p)
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

		return Collector{}, fmt.Errorf("service.AddNewAgentAsSuperCollector: error: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return Collector{}, fmt.Errorf("error commiting txn: error: %w", err)
	}

	return makeServiceCollector(collector)
}

func (s *Service) GetCollector(ctx context.Context, collectorEmail string) (Collector, error) {
	collector, err := s.dbQuery.GetCollector(ctx, collectorEmail)
	if err != nil {
		return Collector{}, err
	}
	return makeServiceCollector(collector)
}
func (s *Service) GetAllCollectors(ctx context.Context) ([]Collector, error) {

	q, err := s.dbQuery.GetAllCollectors(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]Collector, 0, len(q))

	for _, v := range q {
		col, err := makeServiceCollector(v)
		if err != nil {
			continue
		}
		res = append(res, col)
	}
	return slices.Clip(res), nil
}

func (s *Service) GetWasteLogByCollectorPaginated(ctx context.Context, agentID uuid.UUID, limit, offset int32) ([]WasteLog, error) {

	agentId, err := googleUUIDToPgtypeUUID(agentID)
	if err != nil {
		return nil, fmt.Errorf("error formating agentId to pgtype: error: %w", err)
	}

	res, err := s.dbQuery.GetWasteLogByCollectorPaginated(
		ctx,
		postgres.GetWasteLogByCollectorPaginatedParams{
			AgentID: agentId,
			Limit:   limit,
			Offset:  offset,
		},
	)
	if err != nil {
		return nil, err
	}

	r := make([]WasteLog, len(res))
	for i, v := range res {
		wasteID, _ := pgtypeUUIDToGoogleUUID(v.WasteID)
		agentID, _ := pgtypeUUIDToGoogleUUID(v.AgentID)
		r[i] = WasteLog{
			WasteID:              wasteID,
			Claimed:              v.Claimed,
			GlassWeighed:         v.GlassWeighed,
			MetalWeighed:         v.MetalWeighed,
			PlasticsWeighed:      v.PlasticsWeighed,
			PaperWeighed:         v.PaperWeighed,
			GlassPointsEarned:    v.GlassPointsEarned,
			MetalPointsEarned:    v.MetalPointsEarned,
			PlasticsPointsEarned: v.PlasticsPointsEarned,
			PaperPointsEarned:    v.PaperPointsEarned,
			TotalPointsEarned:    v.TotalPointsEarned,
			LocationAddr:         v.LocationAddr,
			TxnName:              v.TxnName,
			StatusReport:         v.StatusReport,
			CollectedRoute:       Agenttype(v.CollectedRoute),
			AgentID:              agentID,
			AccountEmail:         v.AccountEmail,
			CreatedAt:            v.CreatedAt.Time,
		}
	}
	return r, nil
}
func (s *Service) GetWasteByCollectorForMonth(ctx context.Context, agentID uuid.UUID, timespan time.Time) (*WasteByCollectorForMonthResponse, error) {
	agentId, err := googleUUIDToPgtypeUUID(agentID)
	if err != nil {
		return nil, fmt.Errorf("error formatting agentId to pgtype: %w", err)
	}

	res, err := s.dbQuery.GetWasteByCollectorForMonth(
		ctx,
		postgres.GetWasteByCollectorForMonthParams{
			AgentID: agentId,
			Year:    int32(timespan.Year()),
			Month:   int32(timespan.Month()),
		},
	)
	if err != nil {
		return nil, err
	}

	parseJSONToFloatSlice := func(data []byte) []float64 {
		if len(data) == 0 {
			return nil
		}

		var values []float64
		if err := json.Unmarshal(data, &values); err != nil {
			return nil
		}

		return values
	}

	response := &WasteByCollectorForMonthResponse{
		TotalGlass:     toFloat64(res.TotalGlass),
		TotalMetal:     toFloat64(res.TotalMetal),
		TotalPlastics:  toFloat64(res.TotalPlastics),
		TotalPaper:     toFloat64(res.TotalPaper),
		TotalWasteLogs: res.TotalWasteLogs,
		GlassTrend:     parseJSONToFloatSlice(res.GlassTrend),
		MetalTrend:     parseJSONToFloatSlice(res.MetalTrend),
		PlasticTrend:   parseJSONToFloatSlice(res.PlasticTrend),
		PaperTrend:     parseJSONToFloatSlice(res.PaperTrend),
	}

	return response, nil
}
