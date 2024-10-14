package http

import (
	"beltech/service"
	"context"
	"fmt"
)

func (a *App) userFromCache(ctx context.Context, userEmail string) (*service.User, error) {
	x, found := a.cacheStorage.Get(userEmail)
	if found {
		fmt.Println("USER FOUND IN CACHE")
		return x.(*service.User), nil
	}

	fmt.Println("user not found in cache")
	user, err := a.addUserToCache(ctx, userEmail)
	if err != nil {
		fmt.Printf("error adding to cache: %v\n", err)
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
