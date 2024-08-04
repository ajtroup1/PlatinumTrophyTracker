package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/ajtroup1/platinum-trophy-tracker/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store models.UserPlatformAccountStore
}

func NewHandler(store models.UserPlatformAccountStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/accounts/{id:[0-9]+}", h.handleGetUserByID).Methods("GET")
	router.HandleFunc("/update-user-accounts", h.handleUpdateAccounts).Methods("PUT")	
}

func (h *Handler) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: %v", err))
		return
	}

	u, err := h.store.GetAccountsByUserID(uint(id))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, u)
}

func (h *Handler) handleUpdateAccounts(w http.ResponseWriter, r *http.Request) {
	var payload models.UpdateAccountsPayload

	// Decode the request body into the payload struct
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Update user accounts
	err = h.store.UpdateUserAccounts(payload.UserID, payload.Accounts)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update user accounts: %v", err))
		return
	}

	// Send success response
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Successfully updated user accounts"})
}
