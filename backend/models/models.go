package models

import "time"

// USER
type User struct {
	ID             uint32                `json:"id"`
	Username       string                `json:"username"`
	Password       string                `json:"password"`
	Firstname      string                `json:"firstname"`
	Lastname       string                `json:"lastname"`
	Email          string                `json:"email"`
	ImgURL         string                `json:"imgurl"`
	CreatedAt      time.Time             `json:"createdAt"`
	Accounts       []UserPlatformAccount `json:"Accounts"`
	TrackedGames   int                   `json:"trackedGames"`
	CompletedGames int                   `json:"completedGames"`
	LastLogin      time.Time             `json:"lastLogin"`
	Deactivated    bool                  `json:"deactivated"`
}

type UserStore interface {
	GetAllUsers() ([]*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
	EditUser(User) error
	ChangePassword(id uint, currentPassword, newPassword, confirmNewPassword string) error
}

type UserPlatformAccount struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userID"`
	Username   string `json:"username"`
	PlatformID uint32 `json:"platformID"`
}

type UserPlatformAccountStore interface {
	GetAccountsByUserID(id uint) ([]*UserPlatformAccount, error)
	UpdateUserAccounts(userID uint, accounts []*UserPlatformAccount) error
}

type UpdateAccountsPayload struct {
	UserID   uint                   `json:"userID" validate:"required"`
	Accounts []*UserPlatformAccount `json:"accounts" validate:"required"`
}

type RegisterUserPayload struct {
	Username  string `json:"username" validate:"required,min=4,max=25"`
	Password  string `json:"password" validate:"required,password"`
	Firstname string `json:"firstname" validate:"required,min=2,max=255"`
	Lastname  string `json:"lastname" validate:"required,min=2,max=255"`
	Email     string `json:"email" validate:"required,email"`
	ImgLink   string `json:"imglink"`
}

type LoginUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type EditUserPayload struct {
	ID        uint32 `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	ImgURL    string `json:"imgurl"`
}

type ChangePasswordPayload struct {
	UserID             uint   `json:"userID" validate:"required"`
	CurrentPassword    string `json:"currentPassword" validate:"required"`
	NewPassword        string `json:"newPassword" validate:"required,password"`
	ConfirmNewPassword string `json:"confirmPassword" validate:"required"`
}

// GAME
type Game struct {
	ID            uint32     `json:"id"`
	RAWGID        uint       `json:"rawgID"`
	Name          string     `json:"name"`
	Slug          string     `json:"slug"`
	Description   string     `json:"description"`
	Platforms     []Platform `json:"platforms"`
	ReleaseDate   string     `json:"releaseDate"`
	BackgroundIMG string     `json:"backgroundImg"`
	Rating        uint       `json:"rating"`
	Website       string     `json:"website"`
	Genres        []string   `json:"genres"`
	CreatedAt     time.Time  `json:"createdAt"`
}

type GameStore interface {
	GetAllGames() ([]*Game, error) // for dev purposes
	GetGameByID(id uint) (*Game, error)
	AddGamePlatform(name string, gameID uint32) error
	AddGameGenre(name string, gameID uint32) error
	AddGame(game Game) (Game, error)
	AddAchievement(achievement Achievement) (int32, error)
	AddUserAchievement(userID, gameID, achID uint32) error
}

type RAWGGame struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	BackgroundIMG string `json:"background_image"`
}

type RAWGGameResponse struct {
	Results []RAWGGame `json:"results"`
}

type ReturnSearchGamePayload struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	CoverURL string `json:"cover_url"`
}

// Main struct for the JSON response
type GameResponse struct {
	ID                        int               `json:"id"`
	Slug                      string            `json:"slug"`
	Name                      string            `json:"name"`
	NameOriginal              string            `json:"name_original"`
	Description               string            `json:"description_raw"`
	Metacritic                int               `json:"metacritic"`
	MetacriticPlatforms       []interface{}     `json:"metacritic_platforms"`
	Released                  string            `json:"released"`
	TBA                       bool              `json:"tba"`
	Updated                   string            `json:"updated"`
	BackgroundImage           string            `json:"background_image"`
	BackgroundImageAdditional string            `json:"background_image_additional"`
	Website                   string            `json:"website"`
	Rating                    float64           `json:"rating"`
	RatingTop                 int               `json:"rating_top"`
	Reactions                 map[string]int    `json:"reactions"`
	Added                     int               `json:"added"`
	AddedByStatus             map[string]int    `json:"added_by_status"`
	Playtime                  int               `json:"playtime"`
	ScreenshotsCount          int               `json:"screenshots_count"`
	MoviesCount               int               `json:"movies_count"`
	CreatorsCount             int               `json:"creators_count"`
	AchievementsCount         int               `json:"achievements_count"`
	ParentAchievementsCount   int               `json:"parent_achievements_count"`
	RedditURL                 string            `json:"reddit_url"`
	RedditName                string            `json:"reddit_name"`
	RedditDescription         string            `json:"reddit_description"`
	RedditLogo                string            `json:"reddit_logo"`
	RedditCount               int               `json:"reddit_count"`
	TwitchCount               int               `json:"twitch_count"`
	YouTubeCount              int               `json:"youtube_count"`
	ReviewsTextCount          int               `json:"reviews_text_count"`
	RatingsCount              int               `json:"ratings_count"`
	SuggestionsCount          int               `json:"suggestions_count"`
	AlternativeNames          []interface{}     `json:"alternative_names"`
	MetacriticURL             string            `json:"metacritic_url"`
	ParentsCount              int               `json:"parents_count"`
	AdditionsCount            int               `json:"additions_count"`
	GameSeriesCount           int               `json:"game_series_count"`
	UserGame                  interface{}       `json:"user_game"`
	ReviewsCount              int               `json:"reviews_count"`
	SaturatedColor            string            `json:"saturated_color"`
	DominantColor             string            `json:"dominant_color"`
	Platforms                 []PlatformRequest `json:"platforms"`
	Genres                    []Genre           `json:"genres"`
}

// Platform struct for platforms **From response
type PlatformRequest struct {
	Platform   PlatformDetail `json:"platform"`
	ReleasedAt string         `json:"released_at"`
}

// PlatformDetail struct for platform details
type PlatformDetail struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Image           string `json:"image"`
	YearEnd         int    `json:"year_end"`
	YearStart       int    `json:"year_start"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

// Genre struct for genres
type Genre struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	GamesCount      int    `json:"games_count"`
	ImageBackground string `json:"image_background"`
}

// Platform struct for platforms from db
type Platform struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ImgURL      string `json:"imgurl"`
	ReleaseYear string `json:"releaseYear"`
}

// USER GAME
type UserGame struct {
	ID          uint32    `json:"id"`
	UserID      uint32    `json:"userID"`
	GameID      uint32    `json:"gameID"`
	TrackedAt   time.Time `json:"trackedAt"`
	CompletedAt time.Time `json:"completedAt"`
	UpdatedAt   time.Time `json:"updatedAt"` // Altering achievements updates this value
}

type UserGameStore interface {
	GetAllUserGames(userID uint32) ([]*UserGame, error)
	GetUserGameByID(userID, gameID uint32) (*UserGame, error)
	TrackGame(userID, gameID uint32) error
	UntrackGame(userID, gameID uint32) error
}

type TrackGamePayload struct {
	UserID uint32 `json:"userID"`
	GameID uint32 `json:"gameID"`
}

// Achievement
type Achievement struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImgURL      string `json:"image"`
	Percent     string `json:"percent"`
	GameID      uint   `json:"gameID"`
}

type AchievementStore interface {
	CompleteAchievement(userID, achievementID uint32) error
	GetAllAchievementsByGame(gameID uint32) ([]*Achievement, error)
}

type AddAchievementPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imgurl"`
	Percent     float32 `json:"percent"`
}

// USER ACHIEVEMENT
type UserAchievement struct {
	ID            uint32    `json:"id"`
	Completed     bool      `json:"completed"` // Player has unlocked achievement
	UserID        uint32    `json:"userID"`
	GameID        uint32    `json:"gameID"` // References Game.ID NOT UserGame.ID
	AchievementID uint32    `json:"achievementID"`
	CompletedAt   time.Time `json:"completedAt"`
	CreatedAt     time.Time `json:"createdAt"`
}

type UserAchievementStore interface {
	GetAllUserAchievements() ([]*UserAchievement, error)
	GetUserAchievementByID(id int) (*UserAchievement, error)
	CreateUserAchievement(Game) error
}

type CreateUserAchievementPayload struct {
	UserID        uint32 `json:"userID"`
	GameID        uint32 `json:"gameID"` // References Game.ID NOT UserGame.ID
	AchievementID uint32 `json:"achievementID"`
}

type CompletedUserAchievementPayload struct {
	UserID            uint32 `json:"userID"`
	AchievementID uint32 `json:"achievementID"`
}
