package http

import (
	"net/http"
)

func (a *App) authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get the session
		session, err := a.sessionStore.Get(r, "beltech-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the session is new or if it lacks user information
		if session.IsNew || session.Values["userID"] == nil {
			http.Redirect(w, r, "/signup", http.StatusUnauthorized)
			return
		}

		// Session is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
