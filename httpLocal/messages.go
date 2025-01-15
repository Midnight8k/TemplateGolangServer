package httpLocal

import (
	"blankproject/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	SERVICE entities.MessageService
}

func (h *MessageHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getMessage(w, r)
	case http.MethodPut:
		fmt.Printf("To be implemented!")
	case http.MethodPost:
		h.createMessage(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *MessageHandler) getMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Failed to get Messages", http.StatusInternalServerError)
		return
	}

	query := r.URL.Query()
	idParam := query.Get("id")

	if idParam == "" {
		http.Error(w, "Id is required!", http.StatusBadRequest)
		return
	}

	id, _ := strconv.ParseInt(idParam, 10, 64)
	message, err := h.SERVICE.GetMessageById(id)
	if err != nil {
		http.Error(w, "Message could not be found!", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, "Could not create response", http.StatusInternalServerError)
	}
}

func (h *MessageHandler) createMessage(w http.ResponseWriter, r *http.Request) {
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

func (h *MessageHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Failed to get Messages", http.StatusInternalServerError)
		return
	}

	messages, err := h.SERVICE.GetAllMessages()
	if err != nil {
		http.Error(w, "Unable to get messages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Could not create response", http.StatusInternalServerError)
	}
}
