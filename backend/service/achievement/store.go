package achievement

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CompleteAchievement(userID, achievementID uint32) error {
	// Start a transaction to ensure atomicity
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	// Step 1: Complete the Achievement
	_, err = tx.Exec("UPDATE user_achievements SET completed = ?, completed_at = CURRENT_TIMESTAMP WHERE user_id = ? AND achievement_id = ?",
		true, userID, achievementID)
	if err != nil {
		return fmt.Errorf("error updating achievement: %v", err)
	}

	// Step 2: Check if all achievements for the game are completed
	var gameID uint32
	err = tx.QueryRow(`
		SELECT game_id
		FROM user_achievements
		WHERE user_id = ? AND achievement_id = ?`,
		userID, achievementID).Scan(&gameID)
	if err != nil {
		return fmt.Errorf("error retrieving game ID: %v", err)
	}

	// Count completed achievements for the game
	var totalAchievements, completedAchievements int
	err = tx.QueryRow(`
		SELECT COUNT(*)
		FROM achievements
		WHERE game_id = ?`, gameID).Scan(&totalAchievements)
	if err != nil {
		return fmt.Errorf("error counting total achievements: %v", err)
	}

	err = tx.QueryRow(`
		SELECT COUNT(*)
		FROM user_achievements
		WHERE user_id = ? AND game_id = ? AND completed = true`, userID, gameID).Scan(&completedAchievements)
	if err != nil {
		return fmt.Errorf("error counting completed achievements: %v", err)
	}

	// Step 3: Update user_game completed_at if all achievements are completed
	if totalAchievements > 0 && completedAchievements == totalAchievements {
		_, err = tx.Exec(`
			UPDATE user_games
			SET completed_at = ?
			WHERE user_id = ? AND game_id = ?`,
			time.Now(), userID, gameID)
		if err != nil {
			return fmt.Errorf("error updating user_game: %v", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}


func (s *Store) GetAllAchievementsByGame(gameID uint32) ([]*models.Achievement, error) {
	rows, err := s.db.Query("SELECT * FROM achievements WHERE game_id = ?", gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var as []*models.Achievement
	for rows.Next() {
		var a models.Achievement
		err := scanAchievement(rows, &a)
		if err != nil {
			return nil, err
		}
		as = append(as, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return as, nil
}

func scanAchievement(scanner interface {
	Scan(dest ...interface{}) error
}, ach *models.Achievement) error {
	err := scanner.Scan(&ach.ID, &ach.Name, &ach.Description, &ach.ImgURL, &ach.Percent, &ach.GameID)
	if err != nil {
		return err
	}
	return nil
}