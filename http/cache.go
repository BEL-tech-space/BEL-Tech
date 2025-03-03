package http

import (
	"beltech/service"
	"context"
	"fmt"
	"time"
)

const accountCacheDuration = 14 * 24 * time.Hour

func (a *App) addUserToCache(ctx context.Context, email string) error {
	acc, err := a.service.GetAccountByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("addUserToCache: error getting account from db: %w", err)
	}

	if added := a.accountCache.SetWithTTL(
		email,
		acc, 1,
		accountCacheDuration,
	); !added {
		return fmt.Errorf("addUserToCache: error adding account to cache")
	}

	return nil
}

func (a *App) getAccountFromCache(ctx context.Context, accountEmail string) (service.Account, error) {

	acc, found := a.accountCache.Get(accountEmail)
	if found {
		return acc, nil
	}

	if err := a.addUserToCache(ctx, accountEmail); err != nil {
		return service.Account{}, err
	}
	return acc, nil
}

func (a *App) updateAccountCache(acc service.Account) bool {
	return a.accountCache.SetWithTTL(acc.AccountEmail, acc, 1, accountCacheDuration)
}
