package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matheusjost/go-bettings-api/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	bets, err := models.GetAll()
	if err != nil {
		log.Printf("Failed getting all bets: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bets)
}
