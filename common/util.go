package common

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

func ResponseOK(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}

func LoadEnvConfig(t interface{}) {
	if err := envconfig.Process("", t); err != nil {
		log.Fatalf("config: Unable to load config for %T: %s", t, err)
	}
}
