package http

import (
	accountPages "beltech/frontend/pages/account"
	"beltech/service"
)

func account(user service.User) accountPages.User {
	return accountPages.User{
		Email:              user.Email,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		NickName:           user.NickName,
		AvatarUrl:          user.AvatarURL,
		SplWallet:          user.SplWallet,
		CurrBalance:        user.CurrBalance,
		TotalEarned:        user.TotalEarned,
		TotalGlassWaste:    user.TotalGlassWaste,
		TotalMetalWaste:    user.TotalMetalWaste,
		TotalPlasticsWaste: user.TotalPlasticsWaste,
		TotalPaperWaste:    user.TotalPaperWaste,
		LastUpdateEmail:    user.LastUpdateEmail,
	}
}
