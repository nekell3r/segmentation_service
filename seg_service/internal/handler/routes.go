package handler

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/segment/create", h.CreateSegment)
	mux.HandleFunc("/segment/delete", h.DeleteSegment)
	mux.HandleFunc("/segment/rename", h.RenameSegment)
	mux.HandleFunc("/segment/add_user", h.AddUserToSegment)
	mux.HandleFunc("/segment/remove_user", h.RemoveUserFromSegment)
	mux.HandleFunc("/segment/distribute", h.DistributeSegment)
	mux.HandleFunc("/user/segments", h.GetUserSegments)
}
