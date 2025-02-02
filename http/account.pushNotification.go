package http

import (
	accPages "beltech/frontend/pages/account"
	"beltech/service"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/google/uuid"
)

var ActiveNotifications = []service.Notification{}

var newAccountNotifications = []service.Notification{
	{
		NotificationID:      uuid.NewString(),
		NotificationTitle:   "Welcome",
		NotificationHeading: "Welcome to Beltech",
		NotificationBody:    "We are excited to have you on board. We are here to help you make the world a better and greener place.",
		NotificationTime:    time.Now().Add(1 * time.Second).Format(time.RFC3339),
	},
	{
		NotificationID:      uuid.NewString(),
		NotificationTitle:   "Opening Speech",
		NotificationHeading: "Hey Anon, we sent you an email",
		NotificationBody:    "Do well to check your login email, we've shared our opening speech with you.",
		NotificationTime:    time.Now().Add(2 * time.Second).Format(time.RFC3339),
	},
	{
		NotificationID:      uuid.NewString(),
		NotificationTitle:   "Display Name",
		NotificationHeading: "Change your display name",
		NotificationBody:    "Hey Anon,we cannpt both share the same name. You can change your display name to something more personal. Click on your avatar to get started.",
		NotificationTime:    time.Now().Add(3 * time.Second).Format(time.RFC3339),
	},
}

func (a *App) getNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	accountEmail := r.Context().Value(contextKeyEmail).(string)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-AccountId", accountEmail)

	value := r.Header.Get("X-IsNewAccount")
	if v, err := strconv.ParseBool(value); err == nil && v {
		if err := json.NewEncoder(w).Encode(newAccountNotifications); err != nil {
			http.Error(w, "Failed to encode notifications", http.StatusInternalServerError)
			return
		}
		return
	}

	if err := json.NewEncoder(w).Encode(ActiveNotifications); err != nil {
		http.Error(w, "Failed to encode notifications", http.StatusInternalServerError)
	}
}

func (a *App) makeNotificationPageHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	if accountEmail != "giwaoluwatobi@gmail.com" && accountEmail != "kherney" {
		http.Redirect(w, r, "/account/welcome", http.StatusTemporaryRedirect)
		return
	}

	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.MakeNotification(account(user)).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) createNotificationHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	if accountEmail != "giwaoluwatobi@gmail.com" && accountEmail != "kherney" {
		http.Redirect(w, r, "/account/welcome", http.StatusTemporaryRedirect)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	heading := r.FormValue("heading")
	body := r.FormValue("body")
	daysStr := r.FormValue("days")

	if title == "" || heading == "" || body == "" || daysStr == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	daysToLast, err := strconv.ParseInt(daysStr, 10, 32)
	if err != nil {
		http.Error(w, "Days must be a number between 1 and 90", http.StatusBadRequest)
		return
	}

	notification := service.Notification{
		NotificationID:      uuid.NewString(),
		NotificationTitle:   title,
		NotificationHeading: heading,
		NotificationBody:    body,
		NotificationTime:    time.Now().Format(time.RFC3339),
		DaysToLast:          int32(daysToLast),
	}

	if notification.DatabaseId, err = a.service.AddNewNotification(context.Background(), notification.NotificationTitle,
		notification.NotificationHeading, notification.NotificationBody,
		notification.NotificationTime, notification.NotificationID, notification.DaysToLast); err != nil {
		a.serverError(w, err)
		return
	}

	ActiveNotifications = append(ActiveNotifications, notification)
	http.Redirect(w, r, "/account/welcome", http.StatusSeeOther)
}

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
