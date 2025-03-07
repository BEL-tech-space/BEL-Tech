package http

import (
	"fmt"
	"net/http"
)

func (a App) clientError(w http.ResponseWriter, errStatus int, err error) {
	http.Error(w, err.Error(), errStatus)
}

func (a App) serverError(w http.ResponseWriter, err error) {
	fmt.Println("SERVER ERROR:", err.Error())
	// a.logger.Error("serverError" + err.Error())
	a.clientError(w, http.StatusInternalServerError, err)
}
