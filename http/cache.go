package http

import (
	"beltech/service"
	"context"
)

func (a *App) userFromCache(ctx context.Context, userEmail string) (*service.User, error) {
	if a.config.Build != productionEnv {
		return &devServiceUser, nil
	}

	x, found := a.cacheStorage.Get(userEmail)
	if found {
		return x.(*service.User), nil
	}

	user, err := a.addUserToCache(ctx, userEmail)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) addUserToCache(ctx context.Context, email string) (*service.User, error) {
	user, err := a.service.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	a.cacheStorage.SetDefault(email, &user)

	return &user, nil
}
