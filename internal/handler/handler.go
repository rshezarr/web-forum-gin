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

	auth := router.PathPrefix("/user").Subrouter()
	auth.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	auth.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

	post := router.PathPrefix("/post").Subrouter()
	post.HandleFunc("/", h.posts).Methods(http.MethodGet)
	post.HandleFunc("/create", h.createPost).Methods(http.MethodPost)
	post.HandleFunc("/edit", h.editPost).Methods(http.MethodPut)
	post.HandleFunc("/delete", h.deletePost).Methods(http.MethodDelete)
	post.HandleFunc("/like", h.createPost).Methods(http.MethodPost)
	post.HandleFunc("/dislike", h.createPost).Methods(http.MethodPost)

	return router
}
