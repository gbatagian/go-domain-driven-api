package greetme

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GreetMePost() http.HandlerFunc {
	var err error

	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload GreetMeRequestSchema
		err = json.NewDecoder(r.Body).Decode(&requestPayload)
		if err != nil {
			log.Printf("Error decoding request body: %s", err.Error())
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(map[string]string{
			"Greetings": fmt.Sprintf("%s %s", requestPayload.Title, requestPayload.Name),
		})
		if err != nil {
			log.Printf("Error encoding response: %s", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}
