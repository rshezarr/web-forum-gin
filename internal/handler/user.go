package handler

import (
	"encoding/json"
	"forum/internal/model"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user model.User

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("sign up: new decode - %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(data, &user); err != nil {
		logrus.Errorf("sign up: unmarshal - %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := h.service.User.CreateUser(user)
	if err != nil {
		logrus.Errorf("sign up: create user - %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&id); err != nil {
		logrus.Errorf("sign up: response id - %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
