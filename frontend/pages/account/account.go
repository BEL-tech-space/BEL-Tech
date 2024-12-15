package frontend

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Email              string    `json:"email"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	NickName           string    `json:"nick_name"`
	AvatarUrl          string    `json:"avatar_url"`
	SplWallet          string    `json:"spl_wallet"`
	CurrBalance        float64   `json:"curr_balance"`
	TotalEarned        float64   `json:"total_earned"`
	TotalGlassWaste    float32   `json:"total_glass_waste"`
	TotalMetalWaste    float32   `json:"total_metal_waste"`
	TotalPlasticsWaste float32   `json:"total_plastics_waste"`
	TotalPaperWaste    float32   `json:"total_paper_waste"`
	LastUpdateEmail    time.Time `json:"last_update_email"`
}

func toTitle(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func chartUrlGreen(url string) string {
	return fmt.Sprintf(`
	<div
		style="
			background-color: #2FBE85;
			width: 7.5rem;
			height: 3rem;
			mask: url('%s') no-repeat center / contain;
			-webkit-mask: url('%s') no-repeat center / contain;"
	></div>`, url, url)
}

func chartUrlRed(url string) string {
	return fmt.Sprintf(`
	<div
		style="
			background-color: #F6455D;
			width: 7.5rem;
			height: 3rem;
			mask: url('%s') no-repeat center / contain;
			-webkit-mask: url('%s') no-repeat center / contain;"
	></div>`, url, url)
}
