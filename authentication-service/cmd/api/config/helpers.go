package config

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *Config) readJson(writer http.ResponseWriter, request *http.Request, data interface{}) error {
	maxBytes := 1048576

	request.Body = http.MaxBytesReader(writer, request.Body, int64(maxBytes))

	decode := json.NewDecoder(request.Body)
	err := decode.Decode(data)
	if err != nil {
		return err
	}

	err = decode.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body should have one a single JSON")
	}
	return nil
}

func (app *Config) writeJson(writer http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			writer.Header()[key] = value
		}
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, err = writer.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (app *Config) errorJson(writer http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := JsonResponse{
		Error:   true,
		Message: err.Error(),
	}
	return app.writeJson(writer, statusCode, payload)
}
