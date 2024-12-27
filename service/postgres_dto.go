package service

import (
	"beltech/service/internal/repository/postgres"
	"beltech/types"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func makeUpsertUserParams(acc Account) postgres.UpsertAccountParams {
	return postgres.UpsertAccountParams{
		AccountID:          acc.AccountID,
		Email:              acc.Email,
		FirstName:          acc.FirstName,
		LastName:           acc.LastName,
		DisplayName:        acc.DisplayName,
		AvatarUrl:          acc.AvatarUrl,
		AccessToken:        acc.AccessToken,
		AccessTokenExpires: pgtype.Timestamp{Time: acc.AccessTokenExpires, InfinityModifier: pgtype.Finite, Valid: true},
		IDToken:            acc.IDToken,
	}
}

type Account struct {
	Email              string    `json:"email"`
	AccountID          string    `json:"account_id"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	DisplayName        string    `json:"display_name"`
	AvatarUrl          string    `json:"avatar_url"`
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

func pgAccountToserviceAccount(pgAcc postgres.Account) Account {
	return Account{
		AccountID:          pgAcc.AccountID,
		Email:              pgAcc.Email,
		FirstName:          pgAcc.FirstName,
		LastName:           pgAcc.LastName,
		DisplayName:        pgAcc.DisplayName,
		AvatarUrl:          pgAcc.AvatarUrl,
		SplWallet:          pgAcc.SplWallet,
		CurrBalance:        pgAcc.CurrBalance,
		TotalEarned:        pgAcc.TotalEarned,
		TotalGlassWaste:    pgAcc.TotalGlassWaste,
		TotalMetalWaste:    pgAcc.TotalMetalWaste,
		TotalPlasticsWaste: pgAcc.TotalPlasticsWaste,
		TotalPaperWaste:    pgAcc.TotalPaperWaste,
		HashedPassword:     pgAcc.HashedPassword,
		AccessToken:        pgAcc.AccessToken,
		AccessTokenExpires: pgAcc.AccessTokenExpires.Time,
		IDToken:            pgAcc.IDToken,
		LastUpdateEmail:    pgAcc.LastUpdateEmail.Time,
		Role:               pgAcc.Role,
		Joined:             pgAcc.Joined.Time,
		IsDeleted:          pgAcc.IsDeleted,
	}
}

func serviceAccountTopgAccount(acc Account) postgres.Account {
	return postgres.Account{
		AccountID:          acc.AccountID,
		Email:              acc.Email,
		FirstName:          acc.FirstName,
		LastName:           acc.LastName,
		DisplayName:        acc.DisplayName,
		AvatarUrl:          acc.AvatarUrl,
		SplWallet:          acc.SplWallet,
		CurrBalance:        acc.CurrBalance,
		TotalEarned:        acc.TotalEarned,
		TotalGlassWaste:    acc.TotalGlassWaste,
		TotalMetalWaste:    acc.TotalMetalWaste,
		TotalPlasticsWaste: acc.TotalPlasticsWaste,
		TotalPaperWaste:    acc.TotalPaperWaste,
		HashedPassword:     acc.HashedPassword,
		AccessToken:        acc.AccessToken,
		AccessTokenExpires: pgtype.Timestamp{Time: acc.AccessTokenExpires, Valid: true, InfinityModifier: pgtype.Finite},
		IDToken:            acc.IDToken,
		LastUpdateEmail:    pgtype.Timestamp{Time: acc.LastUpdateEmail, Valid: true, InfinityModifier: pgtype.Finite},
		Role:               acc.Role,
		Joined:             pgtype.Timestamp{Time: acc.Joined, Valid: true},
		IsDeleted:          acc.IsDeleted,
	}
}

// type Collector struct {
// 	postgres.Collector
// }

// type GetAllCollectorsEmailAndAgentIDRow struct {
// 	Email   string    `json:"email"`
// 	AgentID uuid.UUID `json:"agent_id"`
// 	Role    postgres.Collectorrole
// }

type AddNewWasteFromCollectorParams struct {
	GlassWeighed    float32     `json:"glass_weighed"`
	MetalWeighed    float32     `json:"metal_weighed"`
	PlasticsWeighed float32     `json:"plastics_weighed"`
	PaperWeighed    float32     `json:"paper_weighed"`
	StMakepoint     interface{} `json:"st_makepoint"`
	StMakepoint_2   interface{} `json:"st_makepoint_2"`
	FunnyName       string      `json:"funny_name"`
	AgentID         uuid.UUID   `json:"agent_id"`
}

// func makeAddNewWasteFromCollectorParams(param AddNewWasteFromCollectorParams) postgres.AddNewWasteFromCollectorParams {
// 	pgUUUID, err := googleUUIDToPgtypeUUID(param.AgentID)
// 	if err != nil {
// 		fmt.Println("error converting google UUID to pgtype UUID:", err)
// 	}
// 	return postgres.AddNewWasteFromCollectorParams{
// 		GlassWeighed:    param.GlassWeighed,
// 		MetalWeighed:    param.MetalWeighed,
// 		PlasticsWeighed: param.PlasticsWeighed,
// 		PaperWeighed:    param.PaperWeighed,
// 		StMakepoint:     param.StMakepoint,
// 		StMakepoint_2:   param.StMakepoint_2,
// 		FunnyName:       param.FunnyName,
// 		AgentID:         pgUUUID,
// 	}
// }

// func broadendUser(u postgres.UpsertAccountRow, acc Account) User {
// 	user.FirstName = u.FirstName
// 	user.LastName = u.LastName
// 	user.DisplayName = u.DisplayName
// 	user.CurrBalance = u.CurrBalance
// 	user.TotalEarned = u.TotalEarned
// 	user.TotalGlassWaste = u.TotalGlassWaste
// 	user.TotalMetalWaste = u.TotalMetalWaste
// 	user.TotalPlasticsWaste = u.TotalPlasticsWaste
// 	user.TotalPaperWaste = u.TotalPaperWaste
// 	return user
// }

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

// func MakeAddNewAgentParams(param AddNewAgentParams) (postgres.AddNewAgentAndCollectorParams, error) {
// 	var r postgres.AddNewAgentAndCollectorParams

// 	switch param.Is {
// 	case "collector":
// 		r.Is = postgres.IdentifyasCollector
// 		if param.Email == "" {
// 			return postgres.AddNewAgentAndCollectorParams{}, errors.New("collector email cannot be empty")
// 		}
// 		r.Email = param.Email
// 		r.Email_2 = param.Email
// 		r.FirstName = param.FirstName
// 		r.LastName = param.LastName
// 		r.NickName = param.NickName
// 		r.AvatarUrl = param.AvatarUrl

// 		if param.Role == "super_collector" {
// 			r.Role = postgres.CollectorroleSuperCollector
// 		} else {
// 			r.Role = postgres.CollectorroleCollector
// 		}

// 	case "bin":
// 		r.Is = postgres.IdentifyasBin

// 	default:
// 		return postgres.AddNewAgentAndCollectorParams{}, errors.New("agent should either be a 'bin' or 'collector'")
// 	}

// 	r.StMakepoint = param.StMakepoint
// 	r.StMakepoint_2 = param.StMakepoint_2
// 	r.City = param.City
// 	r.State = param.State

// 	return r, nil
// }

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

// func makeGetWasteFromCollectorByAgentIDAndDateRow(param []postgres.GetWasteFromCollectorByAgentIDAndDateRow) ([]GetWasteFromCollectorByAgentIDAndDateRow, error) {

// 	var (
// 		r   = make([]GetWasteFromCollectorByAgentIDAndDateRow, 0, len(param))
// 		err error
// 	)

// 	for i := 0; i < len(param); i++ {
// 		r[i].GlassWeighed = param[i].GlassWeighed
// 		r[i].MetalWeighed = param[i].MetalWeighed
// 		r[i].PlasticsWeighed = param[i].PlasticsWeighed
// 		r[i].PaperWeighed = param[i].PaperWeighed
// 		r[i].TempWasteCount = param[i].TempWasteCount
// 		r[i].Location = param[i].Location

// 		if r[i].WasteID, err = pgtypeUUIDToGoogleUUID(param[i].WasteID); err != nil {
// 			return nil, fmt.Errorf("error parsing pgtypeUUID to ggUUID: %w", err)
// 		}
// 		var d any
// 		if d, err = param[i].Date.Value(); err != nil {
// 			return nil, fmt.Errorf("error getting date value: %w", err)
// 		}
// 		r[i].Date = d.(time.Time)
// 	}

// 	return r, nil
// }

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

func pgBeneficiaryToServiceBeneficiary(pgBen postgres.Beneficiary) types.Beneficiary {
	return types.Beneficiary{
		BPhoneNo:        pgBen.BPhoneNo,
		BName:           pgBen.BName,
		ServiceProvider: pgBen.ServiceProvider,
	}
}
