package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListBookings(t *testing.T) {
	rec := httptest.NewRecorder()
	ListBookings(rec, httptest.NewRequest("GET", "/api/bookings", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}
