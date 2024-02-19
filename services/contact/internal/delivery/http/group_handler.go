package http

import (
	"architecture_go/services/contact/internal/domain"
	usecase "architecture_go/services/contact/internal/useCase"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type groupHandler struct {
	groupUseCase usecase.GroupUseCase
}

func NewGroupHandler(groupUseCase usecase.GroupUseCase) *groupHandler {
	return &groupHandler{
		groupUseCase: groupUseCase,
	}
}

func (h *groupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group domain.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if err := h.groupUseCase.CreateGroup(&group); err != nil {
		http.Error(w, fmt.Sprintf("Error creating group: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func (h *groupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	group, err := h.groupUseCase.GetGroup(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving group: %v", err), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(group)
}

func (h *groupHandler) AddContactToGroup(w http.ResponseWriter, r *http.Request) {
	contactIDStr := r.FormValue("contact_id")
	groupIDStr := r.FormValue("group_id")

	contactID, err := strconv.Atoi(contactIDStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	err = h.groupUseCase.AddContactToGroup(contactID, groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Added contact to group")
}
