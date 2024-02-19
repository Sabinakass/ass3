package http

import (
	"architecture_go/services/contact/internal/domain"
	usecase "architecture_go/services/contact/internal/useCase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type contactHandler struct {
	useCase usecase.ContactUseCase
}

func NewContactHandler(useCase usecase.ContactUseCase) *contactHandler {
	return &contactHandler{useCase: useCase}
}

func (h *contactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.useCase.CreateContact(&contact); err != nil {
		http.Error(w, fmt.Sprintf("Error creating contact: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func (h *contactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	contact, err := h.useCase.GetContact(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving contact: %v", err), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func (h *contactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.useCase.UpdateContact(&contact); err != nil {
		http.Error(w, fmt.Sprintf("Error updating contact: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contact)
}

func (h *contactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.useCase.DeleteContact(id); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting contact: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Contact deleted successfully")
}
