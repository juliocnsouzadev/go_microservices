package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

func (app *Config) Broker(writer http.ResponseWriter, request *http.Request) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit Broker Service",
	}

	_ = app.writeJson(writer, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(writer http.ResponseWriter, request *http.Request) {
	var requestPayload RequestPayload
	err := app.readJson(writer, request, &requestPayload)
	if err != nil {
		app.errorJson(writer, err, http.StatusBadRequest)
	}

	switch requestPayload.Action {
	case "auth":
		app.Authenticate(writer, requestPayload.Auth)
	default:
		app.errorJson(writer, errors.New("invalid action"), http.StatusBadRequest)
	}
}

func (app *Config) Authenticate(writer http.ResponseWriter, authPayload AuthPayload) {
	jsonData, _ := json.MarshalIndent(authPayload, "", "\t")
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate",
		bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJson(writer, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJson(writer, err, http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJson(writer, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}
	if response.StatusCode != http.StatusAccepted {
		app.errorJson(writer, errors.New("error auth service"), response.StatusCode)
		return
	}

	var result JsonResponse

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		app.errorJson(writer, err, http.StatusBadRequest)
		return
	}

	if result.Error {
		app.errorJson(writer, errors.New(result.Message), http.StatusUnauthorized)
		return
	}

	resultResponse := JsonResponse{
		Error:   false,
		Message: "Authentication successful",
	}

	app.writeJson(writer, http.StatusAccepted, resultResponse)
}
