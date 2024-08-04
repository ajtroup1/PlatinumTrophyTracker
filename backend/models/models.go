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
	UserID 	   uint	  `json:"userID"`
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
	ID            uint32       `json:"id"`
	Name          string       `json:"name"`
	Slug          string       `json:"slug"`
	Platforms     []Platform   `json:"platforms"`
	ReleaseDate   string       `json:"releaseDate"`
	BackgroundIMG string       `json:"backgroundImg"`
	Rating        float32      `json:"rating"`
	RatingCount   int          `json:"ratingCount"`
	ESRBRating    string       `json:"esrbRating"`
	Genres        []string     `json:"genres"`
	Screenshots   []Screenshot `json:"screenshots"`
	CreatedAt     time.Time    `json:"createdAt"`
}

type GameStore interface {
	GetAllGames() ([]*Game, error)
	GetGameByID(id int) (*Game, error)
	CreateGame(Game) error
	// EditGame(Game) error // don't know if this will exist
}

type AddGamePayload struct {
	Name          string       `json:"name"`
	Slug          string       `json:"slug"`
	Platforms     []Platform   `json:"platforms"`
	ReleaseDate   string       `json:"releaseDate"`
	BackgroundIMG string       `json:"backgroundImg"`
	Rating        float32      `json:"rating"`
	RatingCount   int          `json:"ratingCount"`
	ESRBRating    string       `json:"esrbRating"`
	Genres        []string     `json:"genres"`
	Screenshots   []Screenshot `json:"screenshots"`
}

type Platform struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	ImgURL      string `json:"imgurl"`
	ReleaseYear string `json:"releaseYear"`
}

type Screenshot struct {
	ID     uint32 `json:"id"`
	ImgURL string `json:"imgurl"`
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
	GetAllUserGames() ([]*UserGame, error)
	GetUserGameByID(id int) (*UserGame, error)
	CreateUserGame(UserGame) error
}

// Achievement
type Achievement struct {
	ID          uint32  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imgurl"`
	Percent     float32 `json:"percent"`
}

type AchievementStore interface {
	GetAllAchievements() ([]*Achievement, error)
	GetAchievementByID(id int) (*Achievement, error)
	CreateAchievement(Achievement) error
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
	UserAchievementID uint32 `json:"achievementID"`
}
