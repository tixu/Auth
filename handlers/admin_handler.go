package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tixu/Auth/users"
)

type AdminHandler struct {
	secret    string
	userAdmin users.UserAdmin
}

func (h *AdminHandler) ListAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.listAll(w, r)
	}
}

func (h *AdminHandler) Del() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.del(w, r)
	}
}

func (h *AdminHandler) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.add(w, r)
	}
}

// ListAll the users registered
func (h *AdminHandler) listAll(w http.ResponseWriter, r *http.Request) {

	users, err := h.userAdmin.ListAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Unable to  marshall users", http.StatusServiceUnavailable)
		return
	}
	w.Write(resp)

}

// Delete the user if registered
func (h *AdminHandler) del(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	if err := h.userAdmin.DeleteUser(id); err != nil {
		http.Error(w, "del not implemented", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Add a user to  the users registered
func (h *AdminHandler) add(w http.ResponseWriter, r *http.Request) {
	/*
		var user users.User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		log.Printf("info : %+v\n", user)
	*/

	// h.userAdmin.add(id)
	http.Error(w, "add not implemented", http.StatusServiceUnavailable)
}

// GetAdmin returns the admin stakeholer

// je pense que nous pouvons la remplacer par une fonction qui retourne des fonctions
func GetAdmin(secret string, userService users.UserAdmin) *AdminHandler {

	return &AdminHandler{
		secret:    secret,
		userAdmin: userService,
	}
}
