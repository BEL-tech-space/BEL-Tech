package service

import (
	"beltech/service/internal/repository/postgres"
	"context"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) CreateOrUpdateUser(ctx context.Context, user Account) (Account, error) {
	pgAcc, err := s.dbQuery.UpsertAccount(ctx, makeUpsertUserParams(user))
	if err == nil {
		return pgAccountToserviceAccount(pgAcc), nil
	}
	// handle error
	return Account{}, err
}

func (s *Service) GetAccountByEmail(ctx context.Context, email string) (Account, error) {
	user, err := s.dbQuery.GetAccountByEmail(ctx, email)
	if err != nil {
		fmt.Println("error in get account", err)
		return Account{}, err
	}
	return pgAccountToserviceAccount(user), nil
}

func (s *Service) AddNewBeneficiary(ctx context.Context, bPhoneNo, bName, bSP, email string) ([]Beneficiary, error) {
	bens, err := s.dbQuery.AddNewBeneficiary(ctx,
		postgres.AddNewBeneficiaryParams{
			BPhoneNo:        bPhoneNo,
			BName:           bName,
			ServiceProvider: bSP,
			AccountEmail:    pgtype.Text{String: email, Valid: true},
		},
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			switch pgErr.Code {
			case "23505":
				return nil, fmt.Errorf("%s: %w", pgErr.Message, ErrUniqueConstraintViolation)
			case "23503":
				return nil, fmt.Errorf("%s: %w", pgErr.Message, ErrForeigKeyViolationt)
			default:
				return nil, fmt.Errorf("%s: %w", pgErr.Message, ErrUnaccountedDBError)
			}
		}
		// Handle non-PostgreSQL errors
		return nil, err
	}

	beneficiaries := make([]Beneficiary, 0, len(bens))
	for _, ben := range bens {
		beneficiaries = append(beneficiaries, pgBeneficiaryToServiceBeneficiary(ben))
	}

	return beneficiaries, nil
}

func (s *Service) DeleteBeneficiary(ctx context.Context, bPhoneNo, bSP, email string) ([]Beneficiary, error) {
	res, err := s.dbQuery.DeleteBeneficiary(
		ctx,
		postgres.DeleteBeneficiaryParams{
			BPhoneNo:        bPhoneNo,
			ServiceProvider: bSP,
			AccountEmail:    pgtype.Text{String: email, Valid: true},
		},
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			switch pgErr.Code {
			case "23503":
				return nil, errors.New("foreign key violation: The email does not exist in the users table")
			default:
				return nil, fmt.Errorf("other PostgreSQL error: %s", pgErr.Message)
			}
		}

		// Handle non-PostgreSQL errors
		return nil, fmt.Errorf("other error: %v", err)
	}

	dto := func(param []postgres.DeleteBeneficiaryRow) []Beneficiary {
		r := make([]Beneficiary, 0, len(param))
		for i := 0; i < len(param); i++ {
			r = append(r, Beneficiary{
				BPhoneNo:        param[i].BPhoneNo,
				BName:           param[i].BName,
				ServiceProvider: param[i].ServiceProvider,
			})
		}
		return r
	}

	return dto(res), nil
}

func (s *Service) GetBeneficiaries(ctx context.Context, email string) ([]Beneficiary, error) {
	bens, err := s.dbQuery.GetBeneficiariesByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return nil, err
	}

	beneficiaries := make([]Beneficiary, 0, len(bens))
	for _, ben := range bens {
		// arr := []string{}
		beneficiaries = append(beneficiaries, pgBeneficiaryToServiceBeneficiary(ben))
	}

	return beneficiaries, nil
}

func (s *Service) UpdateAccountDisplayName(ctx context.Context, email, newDisplayName string) (Account, error) {
	pgAcc, err := s.dbQuery.UpdateNickName(ctx,
		postgres.UpdateNickNameParams{AccountEmail: email, DisplayName: newDisplayName})
	if err != nil {
		return Account{}, err
	}

	return pgAccountToserviceAccount(pgAcc), nil
}

func (s *Service) AddNewNotification(ctx context.Context, title, heading, body, time, id string, daysToLast int32) (int32, error) {
	return s.dbQuery.CreateNotification(ctx, postgres.CreateNotificationParams{
		NotifTitle:   title,
		NotifHeading: heading,
		NotifBody:    body,
		NotifTime:    time,
		DaysToLast:   daysToLast,
		NotifID:      id,
	})
}

func (s *Service) GetAllNotifications(ctx context.Context) ([]Notification, error) {
	notifications, err := s.dbQuery.GetAllNotifications(ctx)
	if err != nil {
		return nil, err
	}

	notifToDelete := make([]int32, 0, len(notifications))
	filteredNotifications := make([]Notification, 0, len(notifications))

	for _, n := range notifications {
		notifTime, _ := time.Parse(time.RFC3339, n.NotifTime)
		if time.Now().After(notifTime.AddDate(0, 0, int(n.DaysToLast))) {
			notifToDelete = append(notifToDelete, n.ID)
			continue
		}

		filteredNotifications = append(filteredNotifications, postgresNotificationToServiceNotification(n))
	}

	if err = s.dbQuery.DeleteNotificationsByIDs(ctx, slices.Clip(notifToDelete)); err != nil {
		fmt.Println("error deleting notifications", err)
	}

	return slices.Clip(filteredNotifications), nil
}

func (s *Service) AddNewWasteFromCollector(ctx context.Context, wasteLog WasteLog) error {

	agentId, err := googleUUIDToPgtypeUUID(wasteLog.AgentID)
	if err != nil {
		return fmt.Errorf("err formating agentID for %v; error: %w", wasteLog.AgentID.String(), err)
	}
	arg := postgres.AddNewWasteFromCollectorParams{
		Claimed:              wasteLog.Claimed,
		GlassWeighed:         wasteLog.GlassWeighed,
		MetalWeighed:         wasteLog.MetalWeighed,
		PlasticsWeighed:      wasteLog.PlasticsWeighed,
		PaperWeighed:         wasteLog.PaperWeighed,
		GlassPointsEarned:    wasteLog.GlassPointsEarned,
		MetalPointsEarned:    wasteLog.MetalPointsEarned,
		PlasticsPointsEarned: wasteLog.PlasticsPointsEarned,
		PaperPointsEarned:    wasteLog.PaperPointsEarned,
		TotalPointsEarned:    wasteLog.TotalPointsEarned,
		// StMakepoint: wasteLog.st,
		// TODO: combine ST_ Location
		TxnName:      wasteLog.TxnName,
		StatusReport: wasteLog.StatusReport,
		AgentID:      agentId,
		AccountEmail: wasteLog.AccountEmail,
	}

	tx, err := s.pgxPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.dbQuery.WithTx(tx)
	if err := qtx.AddNewWasteFromCollector(ctx, arg); err != nil {
		return err
	}

	if err := qtx.UpdateCollectorStats(
		ctx,
		postgres.UpdateCollectorStatsParams{
			TotalGlassWaste:    float64(wasteLog.GlassWeighed),
			TotalMetalWaste:    float64(wasteLog.MetalWeighed),
			TotalPlasticsWaste: float64(wasteLog.PlasticsWeighed),
			TotalPaperWaste:    float64(wasteLog.PaperWeighed),
			AgentID:            agentId,
		},
	); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetWasteLogByAccountPaginated(
	ctx context.Context,
	accountEmail string,
	limit, offset int32,
) ([]WasteLog, error) {

	res, err := s.dbQuery.GetWasteLogByAccountPaginated(
		ctx,
		postgres.GetWasteLogByAccountPaginatedParams{
			AccountEmail: accountEmail,
			Limit:        limit,
			Offset:       offset,
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

func (s *Service) GetWasteByAccountForMonth(
	ctx context.Context,
	accountEmail string,
	timespan time.Time,
) ([]WasteLog, error) {
	res, err := s.dbQuery.GetWasteByAccountForMonth(
		ctx,
		postgres.GetWasteByAccountForMonthParams{
			AccountEmail: accountEmail,
			Year:         int32(timespan.Year()),
			Month:        int32(timespan.Month()),
		},
	)
	if err != nil {
		return nil, err
	}

	r := make([]WasteLog, len(res))
	for i, v := range res {

		r[i] = WasteLog{
			GlassWeighed:    v.GlassWeighed,
			MetalWeighed:    v.MetalWeighed,
			PlasticsWeighed: v.PlasticsWeighed,
			PaperWeighed:    v.PaperWeighed,
			CreatedAt:       v.CreatedAt.Time,
		}
	}

	return r, nil
}
