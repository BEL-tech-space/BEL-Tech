package service

import (
	"beltech/service/internal/repository/postgres"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func convertToUpsertUserParams(user User) postgres.UpsertUserParams {
	return postgres.UpsertUserParams{
		UserID:             user.UserID,
		Email:              user.Email,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		NickName:           user.NickName,
		AvatarUrl:          user.AvatarUrl,
		SplWallet:          user.SplWallet,
		CurrBalance:        user.CurrBalance,
		TotalEarned:        user.TotalEarned,
		TotalGlassWaste:    user.TotalGlassWaste,
		TotalMetalWaste:    user.TotalMetalWaste,
		TotalPlasticsWaste: user.TotalPlasticsWaste,
		TotalPaperWaste:    user.TotalPaperWaste,
		HashedPassword:     user.HashedPassword,
		AccessToken:        user.AccessToken,
		AccessTokenExpires: user.AccessTokenExpires,
		IDToken:            user.IDToken,
		LastUpdateEmail:    user.LastUpdateEmail,
	}
}

type User struct {
	postgres.User
	IsAuthenticated bool
}

type Collector struct {
	postgres.Collector
}

type GetAllCollectorsEmailAndAgentIDRow struct {
	Email   string    `json:"email"`
	AgentID uuid.UUID `json:"agent_id"`
	Role    postgres.Collectorrole
}

type AddNewTempWasteFromCollectorParams struct {
	GlassWeighed    float32     `json:"glass_weighed"`
	MetalWeighed    float32     `json:"metal_weighed"`
	PlasticsWeighed float32     `json:"plastics_weighed"`
	PaperWeighed    float32     `json:"paper_weighed"`
	StMakepoint     interface{} `json:"st_makepoint"`
	StMakepoint_2   interface{} `json:"st_makepoint_2"`
	AgentID         uuid.UUID   `json:"agent_id"`
}

func makeAddNewTempWasteFromCollectorParams(param AddNewTempWasteFromCollectorParams) postgres.AddNewTempWasteFromCollectorParams {
	pgUUUID, err := googleUUIDToPgtypeUUID(param.AgentID)
	if err != nil {
		fmt.Println("error converting google UUID to pgtype UUID:", err)
	}
	return postgres.AddNewTempWasteFromCollectorParams{
		GlassWeighed:    param.GlassWeighed,
		MetalWeighed:    param.MetalWeighed,
		PlasticsWeighed: param.PlasticsWeighed,
		PaperWeighed:    param.PaperWeighed,
		StMakepoint:     param.StMakepoint,
		StMakepoint_2:   param.StMakepoint_2,
		AgentID:         pgUUUID,
	}
}

type AddNewAgentParams struct {
	Is            string  `json:"is"` // "bin" or "collector"
	Email         string  `json:"email"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	StMakepoint   float64 `json:"st_makepoint"`
	StMakepoint_2 float64 `json:"st_makepoint_2"`
	Role          string  `json:"role"` // "collector" or "super_collector"
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	NickName      string  `json:"nick_name"`
	AvatarUrl     string  `json:"avatar_url"`
}

func MakeAddNewAgentParams(param AddNewAgentParams) (postgres.AddNewAgentAndCollectorParams, error) {
	var r postgres.AddNewAgentAndCollectorParams

	switch param.Is {
	case "collector":
		r.Is = postgres.IdentifyasCollector
		if param.Email == "" {
			return postgres.AddNewAgentAndCollectorParams{}, errors.New("collector email cannot be empty")
		}
		r.Email = pgtype.Text{String: param.Email, Valid: true}
		r.Email_2 = param.Email
		r.FirstName = pgtype.Text{String: param.FirstName, Valid: true}
		r.LastName = pgtype.Text{String: param.LastName, Valid: true}
		r.NickName = pgtype.Text{String: param.NickName, Valid: true}
		r.AvatarUrl = pgtype.Text{String: param.AvatarUrl, Valid: true}

		if param.Role == "super_collector" {
			r.Role = postgres.CollectorroleSuperCollector
		} else {
			r.Role = postgres.CollectorroleCollector
		}

	case "bin":
		r.Is = postgres.IdentifyasBin

	default:
		return postgres.AddNewAgentAndCollectorParams{}, errors.New("agent should either be a 'bin' or 'collector'")
	}

	r.StMakepoint = param.StMakepoint
	r.StMakepoint_2 = param.StMakepoint_2
	r.City = pgtype.Text{String: param.City, Valid: true}
	r.State = pgtype.Text{String: param.State, Valid: true}

	return r, nil
}

type GetWasteFromCollectorByAgentIDAndDateRow struct {
	WasteID         uuid.UUID   `json:"waste_id"`
	Date            time.Time   `json:"date"`
	GlassWeighed    float32     `json:"glass_weighed"`
	MetalWeighed    float32     `json:"metal_weighed"`
	PlasticsWeighed float32     `json:"plastics_weighed"`
	PaperWeighed    float32     `json:"paper_weighed"`
	Location        interface{} `json:"location"`
	TempWasteCount  int64       `json:"temp_waste_count"`
}

func makeGetWasteFromCollectorByAgentIDAndDateRow(param []postgres.GetWasteFromCollectorByAgentIDAndDateRow) ([]GetWasteFromCollectorByAgentIDAndDateRow, error) {

	var (
		r   = make([]GetWasteFromCollectorByAgentIDAndDateRow, 0, len(param))
		err error
	)

	for i := 0; i < len(param); i++ {
		r[i].GlassWeighed = param[i].GlassWeighed
		r[i].MetalWeighed = param[i].MetalWeighed
		r[i].PlasticsWeighed = param[i].PlasticsWeighed
		r[i].PaperWeighed = param[i].PaperWeighed
		r[i].TempWasteCount = param[i].TempWasteCount
		r[i].Location = param[i].Location

		if r[i].WasteID, err = pgtypeUUIDToGoogleUUID(param[i].WasteID); err != nil {
			return nil, fmt.Errorf("error parsing pgtypeUUID to ggUUID: %w", err)
		}
		var d any
		if d, err = param[i].Date.Value(); err != nil {
			return nil, fmt.Errorf("error getting date value: %w", err)
		}
		r[i].Date = d.(time.Time)
	}

	return r, nil
}

func pgtypeUUIDToGoogleUUID(pgUUID pgtype.UUID) (uuid.UUID, error) {
	pgUUIDStr, err := pgUUID.Value()
	if err != nil {
		return uuid.Nil, err
	}
	uuidStr, ok := pgUUIDStr.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("failed to convert pgtype.UUID to string")
	}

	googleUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, err
	}

	if err := uuid.Validate(googleUUID.String()); err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID: %v", err)
	}

	return googleUUID, nil
}

func googleUUIDToPgtypeUUID(googleUUID uuid.UUID) (pgtype.UUID, error) {
	var pgUUID pgtype.UUID
	if err := pgUUID.Scan(googleUUID.String()); err != nil {
		return pgtype.UUID{}, err
	}
	return pgUUID, nil
}
