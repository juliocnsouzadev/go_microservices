package config

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *Config) Broker(writer http.ResponseWriter, request *http.Request) {

	payload := JsonResponse{
		Error:   false,
		Message: "Hit Broker Service",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(out)

}
