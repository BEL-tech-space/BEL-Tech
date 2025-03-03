package service

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountEmail              string    `json:"account_email"`
	AccountID                 string    `json:"account_id"`
	FirstName                 string    `json:"first_name"`
	LastName                  string    `json:"last_name"`
	SplWallet                 string    `json:"spl_wallet"`
	DisplayName               string    `json:"display_name"`
	AvatarUrl                 string    `json:"avatar_url"`
	TotalBelPoint             float64   `json:"total_bel_point"`
	TotalBelToken             float64   `json:"total_bel_token"`
	BelPointBalance           float64   `json:"bel_point_balance"`
	BelTokenBalance           float64   `json:"bel_token_balance"`
	TotalGlassWaste           float64   `json:"total_glass_waste"`
	TotalMetalWaste           float64   `json:"total_metal_waste"`
	TotalPlasticsWaste        float64   `json:"total_plastics_waste"`
	TotalPaperWaste           float64   `json:"total_paper_waste"`
	TotalGlassPointsEarned    float64   `json:"total_glass_points_earned"`
	TotalMetalPointsEarned    float64   `json:"total_metal_points_earned"`
	TotalPlasticsPointsEarned float64   `json:"total_plastics_points_earned"`
	TotalPaperPointsEarned    float64   `json:"total_paper_points_earned"`
	HashedPassword            string    `json:"hashed_password"`
	HashedPin                 string    `json:"hashed_pin"`
	AccessToken               string    `json:"access_token"`
	AccessTokenExpires        time.Time `json:"access_token_expires"`
	IDToken                   string    `json:"id_token"`
	LastUpdateEmail           time.Time `json:"last_update_email"`
	Role                      string    `json:"role"`
	Joined                    time.Time `json:"joined"`
	IsNew                     bool      `json:"is_new"`
	IsDeleted                 bool      `json:"is_deleted"`
	DeletedAt                 time.Time `json:"deleted_at"`
}

type Beneficiary struct {
	BName           string    `json:"b_name"`
	BPhoneNo        string    `json:"b_phone_no"`
	ServiceProvider string    `json:"service_provider"`
	CreatedAt       time.Time `json:"created_at"`
	AccountEmail    string    `json:"account_email"`
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

type WasteLogFromCollector struct {
	WasteID              uuid.UUID   `json:"waste_id"`
	Claimed              bool        `json:"claimed"`
	GlassWeighed         float32     `json:"glass_weighed"`
	MetalWeighed         float32     `json:"metal_weighed"`
	PlasticsWeighed      float32     `json:"plastics_weighed"`
	PaperWeighed         float32     `json:"paper_weighed"`
	GlassPointsEarned    float32     `json:"glass_points_earned"`
	MetalPointsEarned    float32     `json:"metal_points_earned"`
	PlasticsPointsEarned float32     `json:"plastics_points_earned"`
	PaperPointsEarned    float32     `json:"paper_points_earned"`
	TotalPointsEarned    float64     `json:"total_points_earned"`
	LocationAddr         interface{} `json:"location_addr"`
	TxnName              string      `json:"txn_name"`
	AgentID              uuid.UUID   `json:"agent_id"`
	AccountEmail         string      `json:"account_email"`
	CreatedAt            time.Time   `json:"created_at"`
}

type AgentCollector struct {
	AgentID         uuid.UUID     `json:"agent_id"`
	CollectorEmail  string        `json:"collector_email"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	DisplayName     string        `json:"display_name"`
	AvatarUrl       string        `json:"avatar_url"`
	LocationCity    string        `json:"location_city"`
	LocationState   string        `json:"location_state"`
	LocationCountry string        `json:"location_country"`
	LocationAddr    interface{}   `json:"location_addr"`
	CollectorRole   Collectorrole `json:"collector_role"`
	CreatedAt       time.Time     `json:"created_at"`
	IsDeleted       bool          `json:"is_deleted"`
	DeletedAt       time.Time     `json:"deleted_at"`
}

type Collectorrole string

const (
	CollectorroleCollector      Collectorrole = "collector"
	CollectorroleSuperCollector Collectorrole = "super_collector"
)
