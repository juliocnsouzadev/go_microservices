package config

import (
	"net/http"
)

func (app *Config) Auth(writer http.ResponseWriter, request *http.Request) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit Auth Service",
	}

	_ = app.writeJson(writer, http.StatusOK, payload)
}
