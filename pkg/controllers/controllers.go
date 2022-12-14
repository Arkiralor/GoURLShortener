package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Function to handle Index route.
func Index(w http.ResponseWriter, r *http.Request) {
	inp := r.URL.Query().Get("inp")
	if inp == "" {
		inp = "Hello world."
	}
	log.Printf("INFO: %v", inp)

	resp := map[string]string{
		"message": inp,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
