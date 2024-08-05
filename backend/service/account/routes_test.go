package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/gorilla/mux"
)

func TestAccounts(t *testing.T) {
	accountStore := &mockAccountStore{}
	handler := NewHandler(accountStore)

	t.Run("should fail if account payload is invalid", func(t *testing.T) {
		payload := models.UpdateAccountsPayload{
			UserID: 1,
			// missing accounts
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/update-user-accounts", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/update-user-accounts", handler.handleUpdateAccounts)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("failed with status code %d, received %d. Response body: %s", http.StatusBadRequest, rr.Code, rr.Body.String())
		}
	})

	t.Run("should succeed with valid payload", func(t *testing.T) {
		// Test case with valid payload
		payload := models.UpdateAccountsPayload{
			UserID: 1,
			Accounts: []*models.UserPlatformAccount{
				{
					ID:         1,
					Username:   "testuser1",
					PlatformID: 1,
				},
				{
					ID:         2,
					Username:   "testuser2",
					PlatformID: 2,
				},
			},
		}
		marshal, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPut, "/update-user-accounts", bytes.NewBuffer(marshal))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/update-user-accounts", handler.handleUpdateAccounts)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d. Response body: %s", http.StatusOK, rr.Code, rr.Body.String())
		}
	})
}

type mockAccountStore struct{}

func (s *mockAccountStore) GetAccountsByUserID(id uint) ([]*models.UserPlatformAccount, error) {
	return nil, nil
}
func (s *mockAccountStore) UpdateUserAccounts(userID uint, accounts []*models.UserPlatformAccount) error {
	return nil
}
