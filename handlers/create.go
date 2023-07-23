package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/matheusjost/go-bettings-api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var bet models.Bet

	err := json.NewDecoder(r.Body).Decode(&bet)
	if err != nil {
		log.Printf("Error decoding json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(bet)

	var res map[string]any
	if err == nil {
		res = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Bet successfully inserted. ID: %v", id),
		}
	} else {
		res = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("An error occurred while inserting: %v", err),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
