package frontend

import (
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
