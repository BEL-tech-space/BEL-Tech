package service

import (
	"beltech/service/internal/repository/postgres"
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
		AccountEmail:              pgAcc.AccountEmail,
		AccountID:                 pgAcc.AccountID,
		FirstName:                 pgAcc.FirstName,
		LastName:                  pgAcc.LastName,
		SplWallet:                 pgAcc.SplWallet,
		DisplayName:               pgAcc.DisplayName,
		AvatarUrl:                 pgAcc.AvatarUrl,
		TotalBelPoint:             pgAcc.TotalBelPoint,
		TotalBelToken:             pgAcc.TotalBelToken,
		BelPointBalance:           pgAcc.BelPointBalance,
		BelTokenBalance:           pgAcc.BelTokenBalance,
		TotalGlassWaste:           pgAcc.TotalGlassWaste,
		TotalMetalWaste:           pgAcc.TotalMetalWaste,
		TotalPlasticsWaste:        pgAcc.TotalPlasticsWaste,
		TotalPaperWaste:           pgAcc.TotalPaperWaste,
		TotalGlassPointsEarned:    pgAcc.TotalGlassPointsEarned,
		TotalMetalPointsEarned:    pgAcc.TotalMetalPointsEarned,
		TotalPlasticsPointsEarned: pgAcc.TotalPlasticsPointsEarned,
		TotalPaperPointsEarned:    pgAcc.TotalPaperPointsEarned,
		HashedPassword:            pgAcc.HashedPassword,
		HashedPin:                 pgAcc.HashedPin,
		AccessToken:               pgAcc.AccessToken,
		AccessTokenExpires:        pgAcc.AccessTokenExpires.Time,
		IDToken:                   pgAcc.IDToken,
		LastUpdateEmail:           pgAcc.LastUpdateEmail.Time,
		Role:                      pgAcc.Role,
		Joined:                    pgAcc.Joined.Time,
		IsNew:                     pgAcc.IsNew,
		IsDeleted:                 pgAcc.IsDeleted,
		DeletedAt:                 pgAcc.DeletedAt.Time,
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

func pgtextToString(t pgtype.Text) string {
	v, _ := t.Value()
	return v.(string)
}

type AddNewAgentAsCollectorParams struct {
	CollectorEmail  string        `json:"collector_email"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	DisplayName     string        `json:"display_name"`
	AvatarUrl       string        `json:"avatar_url"`
	LocationCity    string        `json:"location_city"`
	LocationState   string        `json:"location_state"`
	LocationCountry string        `json:"location_country"`
	StMakepoint     interface{}   `json:"st_makepoint"`
	StMakepoint_2   interface{}   `json:"st_makepoint_2"`
	CollectorRole   Collectorrole `json:"collector_role"`
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

func pgBeneficiaryToServiceBeneficiary(pgBen postgres.Beneficiary) Beneficiary {
	return Beneficiary{
		BPhoneNo:        pgBen.BPhoneNo,
		BName:           pgBen.BName,
		ServiceProvider: pgBen.ServiceProvider,
	}
}

type AddNewWasteFromCollectorParams struct {
	Claimed         bool        `json:"claimed"`
	GlassWeighed    float32     `json:"glass_weighed"`
	MetalWeighed    float32     `json:"metal_weighed"`
	PlasticsWeighed float32     `json:"plastics_weighed"`
	PaperWeighed    float32     `json:"paper_weighed"`
	PointsEarned    float64     `json:"points_earned"`
	StMakepoint     interface{} `json:"st_makepoint"`
	StMakepoint_2   interface{} `json:"st_makepoint_2"`
	TxnName         string      `json:"txn_name"`
	AgentID         uuid.UUID   `json:"agent_id"`
	AccountEmail    string      `json:"account_email"`
}
