package config

import (
	"errors"
	"net/http"
)

func (app *Config) AuthHit(writer http.ResponseWriter, request *http.Request) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit Auth Service",
	}

	_ = app.writeJson(writer, http.StatusOK, payload)
}

func (app *Config) Auth(writer http.ResponseWriter, request *http.Request) {

	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJson(writer, request, &payload)
	if err != nil {
		_ = app.errorJson(writer, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(payload.Email)
	if err != nil {
		app.errorJson(writer, errors.New("invalid credentials"), http.StatusBadRequest)
	}

	valid, err := user.PasswordMatches(payload.Password)
	if err != nil || !valid {
		app.errorJson(writer, errors.New("invalid credentials"), http.StatusBadRequest)
	}

	result := JsonResponse{
		Error:   false,
		Message: "Successfully authenticated",
	}

	app.writeJson(writer, http.StatusAccepted, result)
}
