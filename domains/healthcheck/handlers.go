package healthcheck

import (
	"encoding/json"
	"net/http"
)

func HealthCheckGet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{"feeling": "great"})
	}

}
