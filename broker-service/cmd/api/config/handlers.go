package config

import (
	"net/http"
)

func (app *Config) Broker(writer http.ResponseWriter, request *http.Request) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit Broker Service",
	}

	_ = app.writeJson(writer, http.StatusOK, payload)
}
