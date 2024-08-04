package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/ajtroup1/platinum-trophy-tracker/service/auth"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAllUsers() ([]*models.User, error) {
	rows, err := s.db.Query("SELECT id, username, password, firstname, lastname, email, imgurl, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Firstname, &u.Lastname, &u.Email, &u.ImgURL, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("users not found")
	}

	return users, nil
}

func (s *Store) GetUserByUsername(username string) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found with id '%d'", id)
	}

	return u, nil
}

func (s *Store) CreateUser(user models.User) error {
	firstname := capitalizeFirstLetter(user.Firstname)
	lastname := capitalizeFirstLetter(user.Lastname)

	_, err := s.db.Exec("INSERT INTO users (username, password, firstname, lastname, email, imgurl, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.Username, user.Password, firstname, lastname, user.Email, user.ImgURL, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) EditUser(user models.User) error {
	firstname := capitalizeFirstLetter(user.Firstname)
	lastname := capitalizeFirstLetter(user.Lastname)

	_, err := s.db.Exec("UPDATE users SET username = ?, firstname = ?, lastname = ?, email = ?, imgLink = ? WHERE id = ?",
		user.Username, firstname, lastname, user.Email, user.ImgURL, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *Store) ChangePassword(id uint, currentPassword, newPassword, confirmNewPassword string) error {
	if newPassword != confirmNewPassword {
		return fmt.Errorf("new password did not match confirm new password")
	}

	// Retrieve the user
	row := s.db.QueryRow("SELECT password FROM users WHERE id = ?", id)
	var storedPassword string
	err := row.Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found with id '%d'", id)
		}
		return err
	}

	// Check the current password
	if !auth.ComparePasswords(storedPassword, []byte(currentPassword)) {
		return fmt.Errorf("current password did not match")
	}

	// Check if new password is the same as the current one
	if auth.ComparePasswords(storedPassword, []byte(newPassword)) {
		return fmt.Errorf("new password cannot be the same as current password")
	}

	// Hash the new password
	hashedPassword, err := auth.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("error hashing new password: %v", err)
	}

	// Update the password in the database
	_, err = s.db.Exec("UPDATE users SET password = ? WHERE id = ?", hashedPassword, id)
	if err != nil {
		return err
	}

	return nil
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func scanRowsIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Email, &user.ImgURL, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
