package http

import (
	"fmt"
	"net/http"
)

var ErrInvalidCookieValue = fmt.Errorf("securecookie: the value is not valid")

func (a App) clientError(w http.ResponseWriter, errStatus int, err error) {
	http.Error(w, err.Error(), errStatus)
	fmt.Println("CLIENT ERROR:", err.Error())
}

func (a App) serverError(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	fmt.Println("SERVER ERROR:", err.Error())
	// a.logger.Error("serverError" + err.Error())
}
