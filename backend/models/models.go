package models

import "time"

// USER
type User struct {
    ID          uint32    `json:"id"`
    Username    string    `json:"username"`
    Password    string    `json:"password"`
    Firstname   string    `json:"firstname"`
    Lastname    string    `json:"lastname"`
    Email       string    `json:"email"`
    PhoneNumber string    `json:"phonenum"`
    ImgURL      string    `json:"imgurl"`
    CreatedAt   time.Time `json:"createdAt"`
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

// Achievement
type Achievement struct {
    ID          uint32  `json:"id"`
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
