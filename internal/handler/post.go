package handler

import (
	"encoding/json"
	"forum/internal/model"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	data, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var post model.Post
	if err := json.Unmarshal(data, &post); err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := h.service.CreatePost(post)
	if err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&id); err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(ctxKeyUser).(int)

	postId, err := strconv.Atoi(mux.Vars(r)["post_id"])
	if err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	newPost := model.Post{
		ID:      postId,
		UserID:  userId,
		Title:   r.FormValue("post_title"),
		Content: r.FormValue("post_content"),
	}

	id, err := h.service.Post.UpdatePost(newPost)
	if err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&id); err != nil {
		logrus.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) likePost(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) dislikePost(w http.ResponseWriter, r *http.Request) {

}
