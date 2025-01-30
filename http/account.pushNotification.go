package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	Endpoint       string `json:"endpoint"`
	EXpirationTime int64  `json:"expirationTime"`
	Keys           struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

func (a *App) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid subscription data", http.StatusBadRequest)
		return
	}

	fmt.Printf("New subscription:%+v", sub)
	w.WriteHeader(http.StatusOK)
}

func sendNotification(sub Subscription, message string) error {
	subscription := &webpush.Subscription{
		Endpoint: sub.Endpoint,
		Keys: webpush.Keys{
			P256dh: sub.Keys.P256dh,
			Auth:   sub.Keys.Auth,
		},
	}

	resp, err := webpush.SendNotification([]byte(message), subscription, &webpush.Options{
		VAPIDPublicKey:  "BAkdR7kgUR-z04hIZka23Zb1CaNrOjB-ToEpw_jZh2lHwqm8Zxu1yizTpZ8ENoa-2gEL-Y_rNi0NbPUqD_rfDoA",
		VAPIDPrivateKey: "8VLowjusIyf0q9aZ2h8y9G7PmaGCa6cDhuhKw8tha2s",
		TTL:             60,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

type Notification struct {
	NotificationID      string `json:"notificationId"`
	NotificationHeading string `json:"notificationHeading"`
	NotificationTitle   string `json:"notificationTitle"`
	NotificationBody    string `json:"notificationBody"`
	NotificationTime    string `json:"notificationTime"`
}

// notificationPageHandler returns 3 sample notifications
func (a *App) notificationPageHandler(w http.ResponseWriter, r *http.Request) {
	notifications := []Notification{
		{
			NotificationID:      "1",
			NotificationTitle:   "Welcome",
			NotificationHeading: "Welcome, Kherney to BelTech",
			NotificationBody:    "Welcome to BelTech, connect your solana wallet to earn BEL tokens",
			NotificationTime:    time.Now().Add(-2 * time.Hour).Format(time.RFC3339),
		},
		{
			NotificationID:      "2",
			NotificationTitle:   "Connect",
			NotificationHeading: "Connect your wallet to earn BEL tokens",
			NotificationBody:    "Welcome to BelTech, connect your solana wallet to earn BEL tokens",
			NotificationTime:    time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
		},
		{
			NotificationID:      "3",
			NotificationTitle:   "New Device",
			NotificationHeading: "Notice! Your account was logged in on a new device",
			NotificationBody:    "Hello Kherney, your account was logged in on a new device, confirm if it is you.",
			NotificationTime:    time.Now().Format(time.RFC3339),
		},
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-AccountId", accountEmail)

	if err := json.NewEncoder(w).Encode(notifications); err != nil {
		http.Error(w, "Failed to encode notifications", http.StatusInternalServerError)
	}
}
