package service

import (
	"beltech/service/internal/repository/postgres"
	"fmt"

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
	CollectorEmail  string      `json:"collector_email"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	DisplayName     string      `json:"display_name"`
	AvatarUrl       string      `json:"avatar_url"`
	LocationCity    string      `json:"location_city"`
	LocationState   string      `json:"location_state"`
	LocationCountry string      `json:"location_country"`
	LocationAddr    interface{} `json:"location_addr"`
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

func makeServiceCollector(param postgres.Collector) (Collector, error) {
	agentId, err := pgtypeUUIDToGoogleUUID(param.AgentID)
	if err != nil {
		return Collector{}, fmt.Errorf("error creatings service collector: error: %w", err)
	}
	return Collector{
		CollectorEmail:  param.CollectorEmail,
		FirstName:       param.FirstName,
		LastName:        param.LastName,
		DisplayName:     param.DisplayName,
		AvatarUrl:       param.AvatarUrl,
		LocationCity:    param.LocationCity,
		LocationState:   param.LocationState,
		LocationCountry: param.LocationCountry,
		LocationAddr:    param.LocationAddr,
		AgentID:         agentId,
		CollectorRole:   Collectorrole(param.CollectorRole),
	}, nil
}
