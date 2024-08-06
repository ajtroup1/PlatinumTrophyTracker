package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/gorilla/mux"
)

func TestUser(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if payload is invalid", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			// Missing required username
			Password:  "Sample123!",
			Firstname: "Adam",
			Lastname:  "Troup",
			Email:     "adamjtroup@gmail.com",
			ImgLink:   "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
		}

		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, received %d. Response body: %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("should register a user successfully", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			Username:  "adamjtroup",
			Password:  "Sample123!",
			Firstname: "Adam",
			Lastname:  "Troup",
			Email:     "adamjtroup@gmail.com",
			ImgLink:   "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
		}

		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, received %d. Response body: %s", http.StatusCreated, rr.Code, rr.Body.String())
		}
	})

	t.Run("should edit a user", func(t *testing.T) {
		payload := models.EditUserPayload{
			ID:        1,
			Username:  "adamjtroup2",
			Firstname: "Adam",
			Lastname:  "Troup",
			Email:     "adamjtroup@gmail.com",
			ImgURL:    "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/edit-user", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/edit-user", handler.handleEdit)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d. Response body: %s", http.StatusOK, rr.Code, rr.Body.String())
		}
	})

	t.Run("should change a user's password", func(t *testing.T) {
		payload := models.ChangePasswordPayload{
			UserID:             1,
			CurrentPassword:    "Sample123!",
			NewPassword:        "Sample1234!",
			ConfirmNewPassword: "Sample1234!",
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/change-password", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/change-password", handler.handleChangePassword)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("failed with status code %d, received %d. Response body: %s", http.StatusOK, rr.Code, rr.Body.String())
		}
	})
}

type mockUserStore struct{}

func (s *mockUserStore) GetAllUsers() ([]*models.User, error) {
	return nil, fmt.Errorf("users not found")
}

func (s *mockUserStore) GetUserByUsername(username string) (*models.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) GetUserByUsernameOrEmail(username string) (*models.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) GetUserByID(id int) (*models.User, error) {
	if id == 1 {
		return &models.User{
			ID:        1,
			Username:  "adamjtroup",
			Password:  "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname: "Adam",
			Lastname:  "Troup",
			Email:     "adamjtroup@gmail.com",
			ImgURL:    "https://upload.wikimedia.org/wikipedia/en/thumb/2/29/DS2_by_Future.jpg/220px-DS2_by_Future.jpg",
			CreatedAt: time.Date(
				2024,      // year
				time.July, // month
				20,        // day
				4,         // hour
				27,        // minute
				55,        // second
				0,         // nanoseconds
				time.UTC,  // location
			),
		}, nil
	} else if id == 2 {
		return &models.User{
			ID:        2,
			Username:  "friend",
			Password:  "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname: "Sample",
			Lastname:  "Name",
			Email:     "sample@mail.com",
			ImgURL:    "",
			CreatedAt: time.Date(
				2024,      // year
				time.July, // month
				20,        // day
				4,         // hour
				27,        // minute
				55,        // second
				0,         // nanoseconds
				time.UTC,  // location
			),
		}, nil
	} else if id == 3 {
		return &models.User{
			ID:        3,
			Username:  "enemy",
			Password:  "$2a$10$luQ7PyQR0KQeliaN15Y55uMFFPzdwDW8VhjEPvIWfJUizTN4IGps2",
			Firstname: "Sample",
			Lastname:  "Name",
			Email:     "sample2@mail.com",
			ImgURL:    "",
			CreatedAt: time.Date(
				2024,      // year
				time.July, // month
				20,        // day
				4,         // hour
				27,        // minute
				55,        // second
				0,         // nanoseconds
				time.UTC,  // location
			),
		}, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (s *mockUserStore) CreateUser(user models.User) error {
	return nil
}

func (s *mockUserStore) EditUser(user models.User) error {
	return nil
}

func (s *mockUserStore) ChangePassword(id uint, currentPassword, newPassword, confirmNewPassword string) error {
	return nil
}

// Correct usage of time.Date
func ParseTime(timestamp string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
