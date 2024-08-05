package achievement

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CompleteAchievement(userID, achievementID uint32) error {
	_, err := s.db.Exec("UPDATE user_achievements SET completed = ? WHERE user_id = ? AND achievement_id = ?",
		true, userID, achievementID)
	if err != nil {
		return fmt.Errorf("error updating achievement: %v", err)
	}

	return nil
}