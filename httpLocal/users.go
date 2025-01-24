package httpLocal

import (
	"blankproject/entities"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	SERVICE entities.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post entities.User

	// Decode the JSON body into the Post struct
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userEntity := &entities.User{
		Nickname:     post.Nickname,
		Username:     post.Username,
		Email:        post.Email,
		Notes:        post.Notes,
		CreationDate: post.CreationDate,
	}

	if err := h.SERVICE.SaveUser(userEntity); err != nil {
		http.Error(w, "Failed to save greeting", http.StatusInternalServerError)
		return
	}

	// Respond with the created post
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
