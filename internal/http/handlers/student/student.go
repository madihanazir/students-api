package student

import (
	"net/http"
)

// New returns an HTTP handler function for the students API
func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Students API is working"))
	}
}
