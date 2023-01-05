package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) posts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.Post.GetAllPosts()
	if err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&posts); err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) editPost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) dislikePost(w http.ResponseWriter, r *http.Request) {

}
