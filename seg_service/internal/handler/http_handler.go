package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"seg_service/internal/domain"
)

type Handler struct {
	service domain.SegmentService
}

func NewHandler(service domain.SegmentService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] CreateSegment called")
	type req struct {
		Name string `json:"name"`
	}
	var body req
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[HTTP] CreateSegment bad request: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] CreateSegment params: name=%s", body.Name)
	if err := h.service.CreateSegment(body.Name); err != nil {
		log.Printf("[HTTP] CreateSegment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[HTTP] CreateSegment success")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] DeleteSegment called")
	name := r.URL.Query().Get("name")
	if name == "" {
		log.Println("[HTTP] DeleteSegment missing name")
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] DeleteSegment params: name=%s", name)
	if err := h.service.DeleteSegment(name); err != nil {
		log.Printf("[HTTP] DeleteSegment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[HTTP] DeleteSegment success")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RenameSegment(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] RenameSegment called")
	type req struct{ OldName, NewName string }
	var body req
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[HTTP] RenameSegment bad request: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] RenameSegment params: old=%s, new=%s", body.OldName, body.NewName)
	if err := h.service.RenameSegment(body.OldName, body.NewName); err != nil {
		log.Printf("[HTTP] RenameSegment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[HTTP] RenameSegment success")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddUserToSegment(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] AddUserToSegment called")
	type req struct {
		UserID  int64
		Segment string
	}
	var body req
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[HTTP] AddUserToSegment bad request: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] AddUserToSegment params: user_id=%d, segment=%s", body.UserID, body.Segment)
	if err := h.service.AddUserToSegment(body.UserID, body.Segment); err != nil {
		log.Printf("[HTTP] AddUserToSegment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[HTTP] AddUserToSegment success")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RemoveUserFromSegment(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] RemoveUserFromSegment called")
	type req struct {
		UserID  int64
		Segment string
	}
	var body req
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[HTTP] RemoveUserFromSegment bad request: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] RemoveUserFromSegment params: user_id=%d, segment=%s", body.UserID, body.Segment)
	if err := h.service.RemoveUserFromSegment(body.UserID, body.Segment); err != nil {
		log.Printf("[HTTP] RemoveUserFromSegment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[HTTP] RemoveUserFromSegment success")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DistributeSegment(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] DistributeSegment called")
	type req struct {
		Segment string
		Percent float64
	}
	var body req
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[HTTP] DistributeSegment bad request: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] DistributeSegment params: segment=%s, percent=%.2f", body.Segment, body.Percent)
	if err := h.service.DistributeSegmentToPercent(body.Segment, body.Percent); err != nil {
		log.Printf("[HTTP] DistributeSegment error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("[HTTP] DistributeSegment success")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetUserSegments(w http.ResponseWriter, r *http.Request) {
	log.Println("[HTTP] GetUserSegments called")
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		log.Printf("[HTTP] GetUserSegments bad user_id: %v", err)
		http.Error(w, "bad user_id", http.StatusBadRequest)
		return
	}
	log.Printf("[HTTP] GetUserSegments params: user_id=%d", userID)
	segments, err := h.service.GetUserSegments(userID)
	if err != nil {
		log.Printf("[HTTP] GetUserSegments error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("[HTTP] GetUserSegments success: %v", segments)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(segments)
}
