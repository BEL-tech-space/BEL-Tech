package http

import (
	accPages "beltech/frontend/pages/account"
	"beltech/service"
)

func account(user service.User) accPages.User {
	return accPages.User{
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
