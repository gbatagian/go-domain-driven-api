package healthcheck

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthCheckGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(map[string]string{"feeling": "great"})
		if err != nil {
			log.Printf("Error encoding response: %s", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}
