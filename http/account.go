package http

import (
	accPages "beltech/frontend/pages/account"
	"beltech/service"
)

func account(acc service.Account) accPages.Account {
	return accPages.Account{
		AccountID:          acc.AccountID,
		AccountEmail:              acc.AccountEmail,
		FirstName:          acc.FirstName,
		LastName:           acc.LastName,
		DisplayName:        acc.DisplayName,
		AvatarUrl:          acc.AvatarUrl,
		SplWallet:          acc.SplWallet,
		CurrBalance:        acc.CurrBalance,
		TotalEarned:        acc.TotalEarned,
		TotalGlassWaste:    acc.TotalGlassWaste,
		TotalMetalWaste:    acc.TotalMetalWaste,
		TotalPlasticsWaste: acc.TotalPlasticsWaste,
		TotalPaperWaste:    acc.TotalPaperWaste,
		HashedPassword:     acc.HashedPassword,
		AccessToken:        acc.AccessToken,
		AccessTokenExpires: acc.AccessTokenExpires,
		IDToken:            acc.IDToken,
		LastUpdateEmail:    acc.LastUpdateEmail,
		Role:               acc.Role,
		Joined:             acc.Joined,
		IsDeleted:          acc.IsDeleted,
	}
}
