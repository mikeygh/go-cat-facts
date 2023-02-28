package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

func facts(router chi.Router) {
	router.Get("/fact", getFact)
}

func getFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
