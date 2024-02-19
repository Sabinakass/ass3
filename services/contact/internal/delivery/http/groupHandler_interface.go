package http

import "net/http"

type GroupHandler interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)
	AddContactToGroup(w http.ResponseWriter, r *http.Request)
}
