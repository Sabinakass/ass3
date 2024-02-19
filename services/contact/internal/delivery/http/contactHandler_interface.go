package http

import "net/http"

type ContactHandler interface {
	GetContact(w http.ResponseWriter, r *http.Request)
	CreateContact(w http.ResponseWriter, r *http.Request)
	DeleteContact(w http.ResponseWriter, r *http.Request)
	UpdateContact(w http.ResponseWriter, r *http.Request)
}
