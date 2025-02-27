package service

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Email              string    `json:"email"`
	AccountID          string    `json:"account_id"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	DisplayName        string    `json:"display_name"`
	AvatarUrl          string    `json:"avatar_url"`
	SplWallet          string    `json:"spl_wallet"`
	CurrBalance        float64   `json:"curr_balance"`
	TotalEarned        float64   `json:"total_earned"`
	TotalGlassWaste    float32   `json:"total_glass_waste"`
	TotalMetalWaste    float32   `json:"total_metal_waste"`
	TotalPlasticsWaste float32   `json:"total_plastics_waste"`
	TotalPaperWaste    float32   `json:"total_paper_waste"`
	HashedPassword     string    `json:"hashed_password"`
	AccessToken        string    `json:"access_token"`
	AccessTokenExpires time.Time `json:"access_token_expires"`
	IDToken            string    `json:"id_token"`
	LastUpdateEmail    time.Time `json:"last_update_email"`
	Role               string    `json:"role"`
	Joined             time.Time `json:"joined"`
	IsDeleted          bool      `json:"is_deleted"`
	IsNew              bool      `json:"is_new"`
}

type Notification struct {
	NotificationID      string `json:"notificationId"`
	NotificationHeading string `json:"notificationHeading"`
	NotificationTitle   string `json:"notificationTitle"`
	NotificationBody    string `json:"notificationBody"`
	NotificationTime    string `json:"notificationTime"`
	DaysToLast          int32
	DatabaseId          int32
}

type WasteLog struct {
	WasteID         uuid.UUID   `json:"waste_id"`
	CreatedAt       time.Time   `json:"created_at"`
	Claimed         bool        `json:"claimed"`
	GlassWeighed    float32     `json:"glass_weighed"`
	MetalWeighed    float32     `json:"metal_weighed"`
	PlasticsWeighed float32     `json:"plastics_weighed"`
	PaperWeighed    float32     `json:"paper_weighed"`
	LocationAddr    interface{} `json:"location_addr"`
	FunnyName       string      `json:"funny_name"`
	AgentID         uuid.UUID   `json:"agent_id"`
	Email           string      `json:"email"`
}
