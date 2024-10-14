package service

import (
	"beltech/service/internal/repository/postgres"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func makeUpsertUserParams(user User) postgres.UpsertUserParams {
	return postgres.UpsertUserParams{
		UserID:             user.UserID,
		Email:              user.Email,
		FirstName:          pgtype.Text{String: user.FirstName, Valid: true},
		LastName:           pgtype.Text{String: user.LastName, Valid: true},
		NickName:           pgtype.Text{String: user.NickName, Valid: true},
		AvatarUrl:          pgtype.Text{String: user.AvatarURL, Valid: true},
		AccessToken:        pgtype.Text{String: user.AccessToken, Valid: true},
		AccessTokenExpires: pgtype.Timestamp{Time: user.AccessTokenExpires, InfinityModifier: pgtype.Finite, Valid: true},
		IDToken:            pgtype.Text{String: user.IDToken, Valid: true},
	}
}

type User struct {
	UserID             string    `json:"user_id"`
	Email              string    `json:"email"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	NickName           string    `json:"nick_name"`
	AvatarURL          string    `json:"avatar_url"`
	SplWallet          string    `json:"spl_wallet"`
	CurrBalance        float64   `json:"curr_balance"`
	TotalEarned        float64   `json:"total_earned"`
	TotalGlassWaste    float32   `json:"total_glass_waste"`
	TotalMetalWaste    float32   `json:"total_metal_waste"`
	TotalPlasticsWaste float32   `json:"total_plastics_waste"`
	TotalPaperWaste    float32   `json:"total_paper_waste"`
	HashedPassword     string    `json:"hashed_password"`
	AccessToken        string    `json:"access_token"`
	AccessTokenExpires time.Time `json:"access_token_expires"`
	IDToken            string    `json:"id_token"`
	LastUpdateEmail    time.Time `json:"last_update_email"`
	Role               string    `json:"role"`
	Joined             time.Time `json:"joined"`
	IsDeleted          bool      `json:"is_deleted"`
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

func broadendUser(u postgres.UpsertUserRow, user User) User {
	user.FirstName = pgtextToString(u.FirstName)
	user.LastName = pgtextToString(u.LastName)
	user.NickName = pgtextToString(u.NickName)
	user.CurrBalance = u.CurrBalance
	user.TotalEarned = u.TotalEarned
	user.TotalGlassWaste = u.TotalGlassWaste
	user.TotalMetalWaste = u.TotalMetalWaste
	user.TotalPlasticsWaste = u.TotalPlasticsWaste
	user.TotalPaperWaste = u.TotalPaperWaste
	return user
}

func pgtextToString(t pgtype.Text) string {
	v, _ := t.Value()
	return v.(string)
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

func pgUserToServiceUser(pu postgres.User) User {
	return User{
		UserID:             pu.UserID,
		Email:              pu.Email,
		FirstName:          pu.FirstName.String,
		LastName:           pu.LastName.String,
		NickName:           pu.NickName.String,
		AvatarURL:          pu.AvatarUrl.String,
		SplWallet:          pu.SplWallet.String,
		CurrBalance:        pu.CurrBalance,
		TotalEarned:        pu.TotalEarned,
		TotalGlassWaste:    pu.TotalGlassWaste,
		TotalMetalWaste:    pu.TotalMetalWaste,
		TotalPlasticsWaste: pu.TotalPlasticsWaste,
		TotalPaperWaste:    pu.TotalPaperWaste,
		HashedPassword:     pu.HashedPassword.String,
		AccessToken:        pu.AccessToken.String,
		AccessTokenExpires: pu.AccessTokenExpires.Time,
		IDToken:            pu.IDToken.String,
		LastUpdateEmail:    pu.LastUpdateEmail.Time,
		Role:               pu.Role,
		Joined:             pu.Joined.Time,
		IsDeleted:          pu.IsDeleted,
	}
}

func serviceUserToPg(u User) postgres.User {
	return postgres.User{
		UserID:             u.UserID,
		Email:              u.Email,
		FirstName:          pgtype.Text{String: u.FirstName, Valid: true},
		LastName:           pgtype.Text{String: u.LastName, Valid: true},
		NickName:           pgtype.Text{String: u.NickName, Valid: true},
		AvatarUrl:          pgtype.Text{String: u.AvatarURL, Valid: true},
		SplWallet:          pgtype.Text{String: u.SplWallet, Valid: true},
		CurrBalance:        u.CurrBalance,
		TotalEarned:        u.TotalEarned,
		TotalGlassWaste:    u.TotalGlassWaste,
		TotalMetalWaste:    u.TotalMetalWaste,
		TotalPlasticsWaste: u.TotalPlasticsWaste,
		TotalPaperWaste:    u.TotalPaperWaste,
		HashedPassword:     pgtype.Text{String: u.HashedPassword, Valid: true},
		AccessToken:        pgtype.Text{String: u.AccessToken, Valid: true},
		AccessTokenExpires: pgtype.Timestamp{Time: u.AccessTokenExpires, Valid: true, InfinityModifier: pgtype.Finite},
		IDToken:            pgtype.Text{String: u.IDToken, Valid: true},
		LastUpdateEmail:    pgtype.Timestamp{Time: u.LastUpdateEmail, Valid: true, InfinityModifier: pgtype.Finite},
		Role:               u.Role,
		Joined:             pgtype.Timestamp{Time: u.Joined, Valid: true},
		IsDeleted:          u.IsDeleted,
	}
}
