package handler

import (
	"forum/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user", h.signUp).Methods(http.MethodPost)

	return router
}
