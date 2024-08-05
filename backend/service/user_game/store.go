package usergame

import (
	"database/sql"
	"fmt"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}



func (s *Store) GetUserGameByID(userID, gameID uint32) (*models.UserGame, error) {
	row := s.db.QueryRow("SELECT * FROM user_games WHERE user_id = ? AND game_id = ?", userID, gameID)

	var u models.UserGame
	err := scanUserGame(row, &u)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user game not found with user_id '%d' and game_id '%d'", userID, gameID)
		}
		return nil, err
	}

	return &u, nil
}

func (s *Store) GetAllUserGames(userID uint32) ([]*models.UserGame, error) {
	rows, err := s.db.Query("SELECT * FROM user_games WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var us []*models.UserGame
	for rows.Next() {
		var u models.UserGame
		err := scanUserGame(rows, &u)
		if err != nil {
			return nil, err
		}
		us = append(us, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return us, nil
}

func (s *Store) TrackGame(userID, gameID uint32) error {
	_, err := s.db.Exec("INSERT INTO user_games (user_id, game_id) VALUES (?, ?)",
		userID, gameID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UntrackGame(userID, gameID uint32) error {
	_, err := s.db.Exec("DELETE FROM user_games WHERE user_id = ? AND game_id = ?", userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to delete from user_games: %v", err)
	}

	_, err = s.db.Exec("DELETE FROM user_achievements WHERE user_id = ? AND game_id = ?", userID, gameID)
	if err != nil {
		return fmt.Errorf("failed to delete from user_achievements: %v", err)
	}

	return nil
}

func scanUserGame(scanner interface {
	Scan(dest ...interface{}) error
}, game *models.UserGame) error {
	err := scanner.Scan(&game.ID, &game.UserID, &game.GameID, &game.TrackedAt, &game.CompletedAt, &game.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
