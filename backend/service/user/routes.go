package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/ajtroup1/platinum-trophy-tracker/service/auth"
	"github.com/ajtroup1/platinum-trophy-tracker/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store models.UserStore
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.handleGetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", h.handleGetUserByID).Methods("GET")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/edit-user", h.handleEdit).Methods("PUT")
	router.HandleFunc("/change-password", h.handleChangePassword).Methods("PUT")
}

func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	us, err := h.store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, us)
}

func (h *Handler) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert the ID from string to integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user id: %v", err))
		return
	}
	u, err := h.store.GetUserByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	utils.WriteJSON(w, http.StatusOK, u)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user models.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.store.GetUserByUsername(user.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// Uncomment and use the appropriate JWT secret from your config
	// secret := []byte(config.Envs.JWTSecret)
	// secret := []byte("your_jwt_secret")
	// token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": ""})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// fmt.Printf("Received payload: %+v\n", payload)

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Check if user exists
	_, err := h.store.GetUserByUsername(payload.Username)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with username %s already exists", payload.Username))
		return
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	u := models.User{
		Username:  payload.Username,
		Password:  hashedPassword,
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		ImgURL:    payload.ImgLink,
		CreatedAt: time.Now(),
	}

	// Create user in the database
	err = h.store.CreateUser(u)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, u)
}

func (h *Handler) handleEdit(w http.ResponseWriter, r *http.Request) {
	var payload models.EditUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Check if user exists
	existingUser, err := h.store.GetUserByID(int(payload.ID))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.ID))
		return
	}

	// Check if any data has changed
	if existingUser.Username == payload.Username &&
		existingUser.Firstname == payload.Firstname &&
		existingUser.Lastname == payload.Lastname &&
		existingUser.Email == payload.Email &&
		existingUser.ImgURL == payload.ImgURL {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("received information is identical to information in database"))
		return
	}

	// Update user details
	existingUser.Username = payload.Username
	existingUser.Firstname = payload.Firstname
	existingUser.Lastname = payload.Lastname
	existingUser.Email = payload.Email
	existingUser.ImgURL = payload.ImgURL

	err = h.store.EditUser(*existingUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, existingUser)
}

func (h *Handler) handleChangePassword(w http.ResponseWriter, r *http.Request) {
	var payload models.ChangePasswordPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to parse JSON: %w", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		log.Printf("Invalid payload: %v", errors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_, err := h.store.GetUserByID(int(payload.UserID))
	if err != nil {
		log.Printf("User with id %d doesn't exist: %v", payload.UserID, err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with id %d doesn't exist", payload.UserID))
		return
	}

	err = h.store.ChangePassword(uint(payload.UserID), payload.CurrentPassword, payload.NewPassword, payload.ConfirmNewPassword)
	if err != nil {
		log.Printf("User with id %d doesn't exist: %v", payload.UserID, err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error with change password function"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
