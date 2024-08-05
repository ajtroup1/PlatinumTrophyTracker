package usergame

import (
	"fmt"
	"net/http"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/ajtroup1/platinum-trophy-tracker/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store models.UserGameStore
	achStore models.AchievementStore
}

func NewHandler(store models.UserGameStore, achStore models.AchievementStore) *Handler {
	return &Handler{store: store, achStore: achStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/track-game", h.handleTrackGame).Methods("POST")
	router.HandleFunc("/complete-achievement", h.handleCompleteAchievement).Methods("POST")
	
}

func (h *Handler) handleTrackGame(w http.ResponseWriter, r *http.Request) {
	var payload models.TrackGamePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Check if user is already tracking game
	_, err := h.store.GetUserGameByID(payload.UserID, payload.GameID)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user is already tracking game"))
		return
	}
	// If not, track the game
	err = h.store.TrackGame(payload.UserID, payload.GameID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error executing game track: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleCompleteAchievement(w http.ResponseWriter, r *http.Request) {
	var payload models.CompletedUserAchievementPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.achStore.CompleteAchievement(payload.UserID, payload.AchievementID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error completing achievement: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}