package service

import "context"

func (s *Service) CreateOrUpdateUser(ctx context.Context, user User) (User, error) {

	u, err := s.dbQuery.UpsertUser(ctx, makeUpsertUserParams(user))
	if err == nil {
		return broadendUser(u, user), nil
	}

	// handle error
	return User{}, err
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := s.dbQuery.GetUser(ctx, email)
	if err != nil {
		return User{}, err
	}
	return pgUserToServiceUser(user), nil
}
