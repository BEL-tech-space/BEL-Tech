package http

import (
	accPages "beltech/frontend/pages/account"
	"context"
	"fmt"
	"net/http"
	"regexp"
)

func (a *App) beneficiariesPageHandler(w http.ResponseWriter, r *http.Request) {

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	user, err := a.getAccountFromCache(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	result, err := a.service.GetBeneficiaries(context.Background(), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.BeneficiariesPage(user.AvatarUrl, result).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func (a *App) addNewBeneficiary(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("err parsing form data; %s", err.Error()))
		return
	}

	if err := validateNewBeneficiary(r); err != nil {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("err validating form data; %s", err.Error()))
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)

	name := r.Form.Get("name")
	if len(name) > 20 {
		name = name[:19]
	}

	result, err := a.service.AddNewBeneficiary(context.Background(),
		r.Form.Get("number"),
		name,
		r.Form.Get("sp"),
		accountEmail,
	)

	if err != nil {
		a.serverError(w, err)
		return
	}

	if r.Form.Get("from") == "beneficiariesPage" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.BeneficiariesMain(result).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}

		return
	}

	if r.Form.Get("from") == "airtimePage" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := accPages.AirtimeBeneficiaries(result).Render(context.TODO(), w); err != nil {
			a.serverError(w, err)
			return
		}

		return
	}
}

func (a *App) deleteBeneficiary(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("err parsing form data; %s", err.Error()))
		return
	}

	if r.Form.Get("name") == "" {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("name is required"))
		return
	}
	if r.Form.Get("phone") == "" {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("phone is required"))
		return
	}
	if r.Form.Get("network") == "" {
		a.clientError(w, http.StatusBadRequest, fmt.Errorf("network is required"))
		return
	}

	accountEmail := r.Context().Value(contextKeyEmail).(string)
	res, err := a.service.DeleteBeneficiary(context.Background(), r.Form.Get("phone"), r.Form.Get("network"), accountEmail)
	if err != nil {
		a.serverError(w, err)
		return
	}

	// result, err := a.service.GetBeneficiaries(context.Background(), accountEmail)
	// if err != nil {
	// 	a.serverError(w, err)
	// 	return
	// }

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := accPages.BeneficiariesMain(res).Render(context.TODO(), w); err != nil {
		a.serverError(w, err)
		return
	}
}

func validateNewBeneficiary(r *http.Request) error {

	bName := r.Form.Get("name")
	bPhoneNo := r.Form.Get("number")
	ServiceProvider := r.Form.Get("sp")

	if bName == "" || bPhoneNo == "" || ServiceProvider == "" {
		return fmt.Errorf("all fields are required")
	}

	if !validPhoneNoRegexp.MatchString(bPhoneNo) {
		return fmt.Errorf("invalid phone number")
	}

	return nil
}

var validPhoneNoRegexp = regexp.MustCompile(`^0[1-9][0-1]\d{8}`)
