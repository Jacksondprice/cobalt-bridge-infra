package handlers

import (
	"encoding/json"
	"net/http"
)

type Booking struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListBookings(w http.ResponseWriter, _ *http.Request) {
	items := []Booking{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
