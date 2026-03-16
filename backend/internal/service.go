package internal

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Service glues handlers to store.
type Service struct {
	store     Store
	adminPass string
}

func NewService(store Store, adminPass string) *Service {
	return &Service{store: store, adminPass: adminPass}
}

func (s *Service) writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (s *Service) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if req.Password != s.adminPass {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	s.writeJSON(w, http.StatusOK, map[string]string{"token": "admin-token"})
}

func (s *Service) AdminGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Admin-Token") != "admin-token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// user auth
func (s *Service) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	created, err := s.store.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	s.writeJSON(w, http.StatusCreated, created)
}

func (s *Service) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		StreetID string `json:"streetId"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	user, ok := s.store.GetUser(req.StreetID, req.Password)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	s.writeJSON(w, http.StatusOK, map[string]string{"token": "user-" + user.StreetID})
}

// public endpoints
func (s *Service) ListBanners(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, s.store.ListBanners())
}

func (s *Service) PublicStudents(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, s.store.ListStudentsPublic())
}

func (s *Service) ListGrowthRoutes(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	s.writeJSON(w, http.StatusOK, s.store.ListRoutesByStudent(id))
}

func (s *Service) CreateGrowthRoute(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req GrowthRoute
	json.NewDecoder(r.Body).Decode(&req)
	req.StudentID = id
	route, err := s.store.CreateRoute(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.writeJSON(w, http.StatusCreated, route)
}

func (s *Service) ReplyRoute(w http.ResponseWriter, r *http.Request) {
	rid := chi.URLParam(r, "rid")
	var req Reply
	json.NewDecoder(r.Body).Decode(&req)
	resp, err := s.store.AddReply(rid, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.writeJSON(w, http.StatusCreated, resp)
}

func (s *Service) ListEvents(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, s.store.ListEvents())
}

func (s *Service) GetEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	s.writeJSON(w, http.StatusOK, s.store.GetEvent(id))
}

// admin endpoints
func (s *Service) AdminStats(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, s.store.Stats())
}

func (s *Service) ListStudents(w http.ResponseWriter, r *http.Request) {
	s.writeJSON(w, http.StatusOK, s.store.ListStudentsAdmin())
}

func (s *Service) PenalizeStudent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var req struct {
		Reason string `json:"reason"`
		Show   bool   `json:"show"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	if err := s.store.Penalize(id, req.Reason, req.Show); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
