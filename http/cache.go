package http

import (
	"beltech/service"
	"context"
	"time"
)

func (a *App) getAccountFromCache(ctx context.Context, accountEmail string) (service.Account, error) {

	acc, found := a.accountCache.Get(accountEmail)
	if found {
		return acc, nil
	}

	acc, err := a.addUserToCache(ctx, accountEmail)
	if err != nil {
		return service.Account{}, err
	}
	return acc, nil
}

func (a *App) addUserToCache(ctx context.Context, email string) (service.Account, error) {
	acc, err := a.service.GetUserByEmail(ctx, email)
	if err != nil {
		return service.Account{}, err
	}

	a.accountCache.SetWithTTL(email, acc, 1, time.Hour)
	return acc, nil
}

func (a *App) updateAccountCache(ctx context.Context, acc service.Account) bool {
	return a.accountCache.SetWithTTL(acc.AccountEmail, acc, 1, time.Hour)
}
