package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KCFLEX/astro-service2.0/repository/entity"
)

type Repository interface {
	GetFixtures(ctx context.Context) ([]entity.Fixtures, error)
	GetFixturesByID(ctx context.Context, id int) (entity.Fixtures, error)
}

type Handler struct {
	repo       Repository
	router     *http.ServeMux
	serverport string
}

func New(repo Repository, mux *http.ServeMux, serverport string) *Handler {
	h := &Handler{
		repo:       repo,
		router:     mux,
		serverport: serverport,
	}
	h.RegisterRoutes()
	return h
}

func (h *Handler) Serve() error {
	//h.RegisterRoutes()
	err := http.ListenAndServe(h.serverport, h.router)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (h *Handler) RegisterRoutes() error {
	h.router.HandleFunc("GET /fixtures", h.GetFixtures)
	h.router.HandleFunc("GET /fixtures/{id}", h.GetFixturesByID)
	return nil
}

func (h *Handler) GetFixtures(w http.ResponseWriter, r *http.Request) {

	getFixtures, err := h.repo.GetFixtures(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve fixtures", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(getFixtures); err != nil {
		http.Error(w, "failed to encode reponse", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GetFixturesByID(w http.ResponseWriter, r *http.Request) {
	parameters := r.PathValue("id")
	fixturesID, err := strconv.Atoi(parameters)
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid fixtures id", http.StatusBadRequest)
		return
	}

	fixtures, err := h.repo.GetFixturesByID(r.Context(), fixturesID)
	if err != nil {
		http.Error(w, "Failed to retrieve fixtures from db", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(fixtures)
	if err != nil {
		http.Error(w, "failed to encode fixtures to response", http.StatusInternalServerError)
	}
}
