package account

import (
	"database/sql"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAccountsByUserID(id uint) ([]*models.UserPlatformAccount, error) {
	rows, err := s.db.Query("SELECT * FROM accounts WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []*models.UserPlatformAccount

	err = scanRowsIntoAccounts(rows, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *Store) CreateAccount(userID uint, account models.UserPlatformAccount) error {
	_, err := s.db.Exec("INSERT INTO accounts (user_id, username, platform_id) VALUES (?, ?, ?)",
		userID, account.Username, account.PlatformID)
	if err != nil {
		return err
	}

	return nil
}


func (s *Store) UpdateUserAccounts(userID uint, accounts []*models.UserPlatformAccount) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM accounts WHERE user_id = ?", userID)
	if err != nil {
		return err
	}

	for _, account := range accounts {
		_, err = tx.Exec("INSERT INTO accounts (user_id, username, platform_id) VALUES (?, ?, ?)",
			userID, account.Username, account.PlatformID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func scanRowsIntoAccounts(rows *sql.Rows, accounts *[]*models.UserPlatformAccount) error {
	for rows.Next() {
		var account models.UserPlatformAccount
		err := rows.Scan(&account.ID, &account.UserID, &account.Username, &account.PlatformID)
		if err != nil {
			return err
		}
		*accounts = append(*accounts, &account)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
