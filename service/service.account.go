package service

import (
	"beltech/service/internal/repository/postgres"
	"beltech/types"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) CreateOrUpdateUser(ctx context.Context, user User) (User, error) {
	u, err := s.dbQuery.UpsertAccount(ctx, makeUpsertUserParams(user))
	if err == nil {
		return broadendUser(u, user), nil
	}
	// handle error
	return User{}, err
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := s.dbQuery.GetAccount(ctx, email)
	if err != nil {
		return User{}, err
	}
	return pgUserToServiceUser(user), nil
}

func (s *Service) AddNewBeneficiary(ctx context.Context, bPhoneNo, bName, bSP, email string) error {
	err := s.dbQuery.AddNewBeneficiary(ctx,
		postgres.AddNewBeneficiaryParams{
			BPhoneNo:        bPhoneNo,
			BName:           bName,
			ServiceProvider: bSP,
			Email:           pgtype.Text{String: email, Valid: true},
		},
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			switch pgErr.Code {
			case "23505":
				return fmt.Errorf("%s: %w", pgErr.Message, ErrUniqueConstraintViolation)
			case "23503":
				return fmt.Errorf("%s: %w", pgErr.Message, ErrForeigKeyViolationt)
			default:
				return fmt.Errorf("%s: %w", pgErr.Message, ErrUnaccountedDBError)
			}
		}
		// Handle non-PostgreSQL errors
		return err
	}

	return nil
}

func (s *Service) DeleteBeneficiary(ctx context.Context, bPhoneNo, bSP, email string) error {
	err := s.dbQuery.DeleteBeneficiary(
		ctx,
		postgres.DeleteBeneficiaryParams{
			BPhoneNo:        bPhoneNo,
			ServiceProvider: bSP,
			Email:           pgtype.Text{String: email, Valid: true},
		},
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			switch pgErr.Code {
			case "23503":
				return errors.New("foreign key violation: The email does not exist in the users table")
			default:
				return fmt.Errorf("other PostgreSQL error: %s", pgErr.Message)
			}
		}

		// Handle non-PostgreSQL errors
		return fmt.Errorf("other error: %v", err)
	}

	return nil
}

func (s *Service) GetBeneficiaries(ctx context.Context, email string) ([]types.Beneficiary, error) {
	bens, err := s.dbQuery.GetBeneficiariesByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return nil, err
	}

	beneficiaries := make([]types.Beneficiary, 0, len(bens))
	for _, ben := range bens {
		beneficiaries = append(beneficiaries, pgBeneficiaryToServiceBeneficiary(ben))
	}

	return beneficiaries, nil
}
