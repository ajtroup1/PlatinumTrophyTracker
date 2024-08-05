package game

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ajtroup1/platinum-trophy-tracker/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllGames() ([]*models.Game, error) {
	rows, err := s.db.Query("SELECT * FROM games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []*models.Game
	for rows.Next() {
		var game models.Game
		err := scanGame(rows, &game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

func (s *Store) GetGameByID(id uint) (*models.Game, error) {
	row := s.db.QueryRow("SELECT * FROM games WHERE id = ?", id)
	var game models.Game
	err := scanGame(row, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (s *Store) AddGame(game models.Game) (models.Game, error) {
	result, err := s.db.Exec("INSERT INTO games (rawg_id, name, slug, description, release_date, background_img, rating, website, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		game.RAWGID, game.Name, game.Slug, game.Description, game.ReleaseDate, game.BackgroundIMG, game.Rating, game.Website, game.CreatedAt)
	if err != nil {
		return models.Game{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return models.Game{}, err
	}

	var insertedGame models.Game
	err = s.db.QueryRow("SELECT id, rawg_id, name, slug, description, release_date, background_img, rating, website, created_at FROM games WHERE id = ?", lastID).Scan(
		&insertedGame.ID, &insertedGame.RAWGID, &insertedGame.Name, &insertedGame.Slug, &insertedGame.Description,
		&insertedGame.ReleaseDate, &insertedGame.BackgroundIMG, &insertedGame.Rating, &insertedGame.Website, &insertedGame.CreatedAt,
	)
	if err != nil {
		return models.Game{}, err
	}

	return insertedGame, nil
}


func (s *Store) AddGamePlatform(name string, gameID uint32) error {
	var platformID uint

	row := s.db.QueryRow("SELECT id FROM platforms WHERE name = ?", name)

	err := row.Scan(&platformID)
	if err != nil {
		if err == sql.ErrNoRows {
			// No platform found with the given name
			return fmt.Errorf("platform with name %s not found", name)
		}
		return err
	}

	_, err = s.db.Exec("INSERT INTO game_platforms (game_id, platform_id) VALUES (?, ?)",
		gameID, platformID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) AddGameGenre(name string, gameID uint32) error {
	_, err := s.db.Exec("INSERT INTO game_genres (game_id, genre) VALUES (?, ?)",
		gameID, name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) AddAchievement(achievement models.Achievement) error {
	_, err := s.db.Exec("INSERT INTO achievements (name, description, imgurl, percent, game_id) VALUES (?, ?, ?, ?, ?)",
		achievement.Name, achievement.Description, achievement.ImgURL, achievement.Percent, achievement.GameID)
	if err != nil {
		return err
	}

	return nil
}


func scanGame(scanner interface {
	Scan(dest ...interface{}) error
}, game *models.Game) error {
	var platforms, genres, screenshots string

	err := scanner.Scan(&game.ID, &game.Name, &game.Slug, &platforms, &game.ReleaseDate, &game.BackgroundIMG, &game.Rating, &genres, &screenshots, &game.CreatedAt)
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(platforms), &game.Platforms); err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(genres), &game.Genres); err != nil {
		return err
	}

	return nil
}
