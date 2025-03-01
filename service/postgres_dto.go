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
		AccountEmail:       acc.AccountEmail,
		FirstName:          acc.FirstName,
		LastName:           acc.LastName,
		DisplayName:        acc.DisplayName,
		AvatarUrl:          acc.AvatarUrl,
		AccessToken:        acc.AccessToken,
		AccessTokenExpires: pgtype.Timestamp{Time: acc.AccessTokenExpires, InfinityModifier: pgtype.Finite, Valid: true},
		IDToken:            acc.IDToken,
	}
}

func pgAccountToserviceAccount(pgAcc postgres.Account) Account {
	return Account{
		AccountID:          pgAcc.AccountID,
		AccountEmail:       pgAcc.AccountEmail,
		FirstName:          pgAcc.FirstName,
		LastName:           pgAcc.LastName,
		DisplayName:        pgAcc.DisplayName,
		AvatarUrl:          pgAcc.AvatarUrl,
		SplWallet:          pgAcc.SplWallet,
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
		IsNew:              pgAcc.IsNew,
	}
}

func serviceAccountTopgAccount(acc Account) postgres.Account {
	return postgres.Account{
		AccountID:          acc.AccountID,
		AccountEmail:       acc.AccountEmail,
		FirstName:          acc.FirstName,
		LastName:           acc.LastName,
		DisplayName:        acc.DisplayName,
		AvatarUrl:          acc.AvatarUrl,
		SplWallet:          acc.SplWallet,
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

func postgresNotificationToServiceNotification(pgNotif postgres.Notification) Notification {
	return Notification{
		NotificationID:      pgNotif.NotifID,
		NotificationHeading: pgNotif.NotifHeading,
		NotificationTitle:   pgNotif.NotifTitle,
		NotificationBody:    pgNotif.NotifBody,
		NotificationTime:    pgNotif.NotifTime,
		DaysToLast:          pgNotif.DaysToLast,
		DatabaseId:          pgNotif.ID,
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

type AddNewAgentAsCollectorParams struct {
	AgentType      Identifyas    `json:"agent_type"`
	City           string        `json:"city"`
	LocationState  string        `json:"location_state"`
	StMakepoint    interface{}   `json:"st_makepoint"`
	StMakepoint_2  interface{}   `json:"st_makepoint_2"`
	CollectorEmail string        `json:"collector_email"`
	CollectorRole  Collectorrole `json:"collector_role"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	DisplayName    string        `json:"display_name"`
	AvatarUrl      string        `json:"avatar_url"`
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

func convertWasteLogToPostgresParams(wasteLog WasteLog) (postgres.AddNewWasteFromCollectorParams, error) {
	agentId, err := googleUUIDToPgtypeUUID(wasteLog.AgentID)
	if err != nil {
		return postgres.AddNewWasteFromCollectorParams{}, err
	}

	return postgres.AddNewWasteFromCollectorParams{
		GlassWeighed:    wasteLog.GlassWeighed,
		MetalWeighed:    wasteLog.MetalWeighed,
		PlasticsWeighed: wasteLog.PlasticsWeighed,
		PaperWeighed:    wasteLog.PaperWeighed,
		TxnName:         wasteLog.TxnName,
		AgentID:         agentId,
		AccountEmail:    pgtype.Text{String: wasteLog.AccountEmail, Valid: true},
		PointsEarned:    wasteLog.PointsEarned,
	}, nil
}
