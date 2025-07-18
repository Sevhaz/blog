package handlers

import (
	"library/models"
	"library/services"
	"library/utils"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.Register(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login models.User
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	user, err := h.Service.GetUserInfo(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var userUpdate models.User
	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]

	err = h.Service.UpdateUser(&userUpdate, userID, claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userUpdate)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	userID := vars["userID"]

	err := h.Service.DeleteUser(claims, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
