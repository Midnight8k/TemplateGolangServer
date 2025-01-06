package httpLocal

import (
	"blankproject/entities"
	"encoding/json"
	"net/http"
)

type MessageHandler struct {
	SERVICE entities.MessageService
}

func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post entities.Message

	// Decode the JSON body into the Post struct
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	messageEntity := &entities.Message{
		Sender:  post.Sender,
		Message: post.Message,
	}

	if err := h.SERVICE.SaveMessage(messageEntity); err != nil {
		http.Error(w, "Failed to save message", http.StatusInternalServerError)
		return
	}

	// Respond with the created post
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
