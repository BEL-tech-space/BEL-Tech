package service

import (
	"beltech/service/internal/repository/postgres"
)


func convertToUpsertUserParams(user User) postgres.UpsertUserParams {
    return postgres.UpsertUserParams{
        UserID:             user.UserID,
        Email:              user.Email,
        FirstName:          user.FirstName,
        LastName:           user.LastName,
        NickName:           user.NickName,
        AvatarUrl:          user.AvatarUrl,
        SplWallet:          user.SplWallet,
        CurrBalance:        user.CurrBalance,
        TotalEarned:        user.TotalEarned,
        TotalGlassWaste:    user.TotalGlassWaste,
        TotalMetalWaste:    user.TotalMetalWaste,
        TotalPlasticsWaste: user.TotalPlasticsWaste,
        TotalPaperWaste:    user.TotalPaperWaste,
        HashedPassword:     user.HashedPassword,
        AcessToken:         user.AcessToken,
        AccessTokenExpires: user.AccessTokenExpires,
        IDToken:            user.IDToken,
        Joined:             user.Joined,
        LastUpdateEmail:    user.LastUpdateEmail,
    }
}