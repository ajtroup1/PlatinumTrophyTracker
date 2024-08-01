# PlatinumTrophyTracker

## API Documentation

### User

- User struct:
  ```go
      type User struct {
          ID          uint32    `json:"id"`
          Username    string    `json:"username"`
          Password    string    `json:"password"`
          Firstname   string    `json:"firstname"`
          Lastname    string    `json:"lastname"`
          Email       string    `json:"email"`
          ImgURL      string    `json:"imgurl"`
          CreatedAt   time.Time `json:"createdAt"`
          Accounts []UserPlatformAccount `json:"Accounts"`
          TrackedGames int `json:"trackedGames"`
          CompletedGames int `json:"completedGames"`
          LastLogin time.Time `json:"lastLogin"`
      }
  ```
- Returns all users
  - Endpoint: `/users`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and []User upon sucessful execution

- Returns user by id
  - Endpoint: `/users/{id}`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and User upon sucessful execution

- Registers a user
  - Endpoint: `/register`
  - Method: `POST`
  - Expects a payload:
     ```go
        type RegisterUserPayload struct {
            Username  string                `json:"username" validate:"required,min=4,max=25"`
            Password  string                `json:"password" validate:"required,password"`
            Firstname string                `json:"firstname" validate:"required,min=2,max=255"`
            Lastname  string                `json:"lastname" validate:"required,min=2,max=255"`
            Email     string                `json:"email" validate:"required,email"`
            ImgLink   string                `json:"imglink"`
            Accounts  []UserPlatformAccount `json:"accounts"`
        }
    ```
  - Returns a 201 upon sucessful execution

- Log in
    - Endpoint: `/login`
    - Method: `POST`
    - Expects a payload:
        ```go
        type LoginUserPayload struct {
            Email    string `json:"email" validate:"required,email"`
            Password string `json:"password" validate:"required"`
        }
        ```
    - Returns a 200 upon successful execution

- Edit user information
    - Endpoint: `/edit-user`
    - Method: `PUT`
    - Expects a payload:
        ```go   
        type EditUser struct {
            Username       string                `json:"username"`
            Firstname      string                `json:"firstname"`
            Lastname       string                `json:"lastname"`
            Email          string                `json:"email"`
            ImgURL         string                `json:"imgurl"`
            Accounts       []UserPlatformAccount `json:"Accounts"`
        }
        ```
    ` Returns a 200 upon successful execution

- Change a user's password
    - Endpoint: `/change-password`
    - Method: `PUT`
    - Expects a payload:
        ```go
        type ChangePasswordPayload struct {
            UserID              uint   `json:"userID" validate:"required"`
            CurrentPassword     string `json:"currentPassword" validate:"required"`
            NewPassword         string `json:"newPassword" validate:"required,password"`
            ConfirmNewPassword  string `json:"confirmPassword" validate:"required"`
        }
        ```
    - Returns 200 upon successful execution

- Deactivate a user's account
    - Endpoint: `/deactivate-account`
    - Method: `DELETE`
    - Expects a payload:
        ```go
        type DeactivateAccountPayload struct {
            UserID uint
            ConfirmPassword string
        }
        ```
    - Returns 200 upon successful execution

### Game

- Game struct:
  ```go
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
  ```

- Returns all games
  - Endpoint: `/games`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and []Game upon sucessful execution

- Returns user by id
  - Endpoint: `/games/{id}`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and Game upon sucessful execution

- Add a game to the db
  - Endpoint: `/add-game`
  - Method: `POST`
  - Expects a payload:
    ```go
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
    ```
    - Returns a 201 upon successful execution

### Achievement

- Achievement struct:
    ```go
    type Achievement struct {
        ID          uint32  `json:"id"`
        Name        string  `json:"name"`
        Description string  `json:"description"`
        ImgURL      string  `json:"imgurl"`
        Percent     float32 `json:"percent"`
    }
    ```

- Returns all achievements
  - Endpoint: `/achievements`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and []Achievement upon sucessful execution

- Returns user by id
  - Endpoint: `/achievement/{id}`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and Achievement upon sucessful execution

- Add an achievement to the db
  -Endpoint: `/add-achievement`
  - Method: `POST`
  - Expects a payload:
    ```go
    type AddAchievementPayload struct {
        Name        string  `json:"name"`
        Description string  `json:"description"`
        ImgURL      string  `json:"imgurl"`
        Percent     float32 `json:"percent"`
    }
    ```
    - Returns a 201 upon successful execution

### User Achievement

- UserAchievement struct:
    ```go
    type UserAchievement struct {
        ID            uint32    `json:"id"`
        Completed     bool      `json:"completed"` // Player has unlocked achievement
        UserID        uint32    `json:"userID"`
        GameID        uint32    `json:"gameID"` // References Game.ID NOT UserGame.ID
        AchievementID uint32    `json:"achievementID"`
        CompletedAt   time.Time `json:"completedAt"`
        CreatedAt     time.Time `json:"createdAt"`
    }
    ```

- Returns all user achievements
  - Endpoint: `/user-achievements`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and []UserAchievement upon sucessful execution

- Returns user achievement by id
  - Endpoint: `/user-achievement/{id}`
  - Method: `GET`
  - Expects no payload
  - Returns a 200 and UserAchievement upon sucessful execution

- Add a user achievement to the db
  - Endpoint: `/add-user-achievement`
  - Method: `POST`
  - Expects a payload:
    ```go
    type CreateUserAchievementPayload struct {
        UserID        uint32    `json:"userID"`
        GameID        uint32    `json:"gameID"` // References Game.ID NOT UserGame.ID
        AchievementID uint32    `json:"achievementID"`
    }
    ```
    - Returns a 201 upon successful execution

- Complete an achievement
  - Endpoint: `completed-achievement`
  - Method: `POST`
  - Expects a payload:
    ```go
    type CompletedUserAchievementPayload struct {
        UserID        uint32    `json:"userID"`
        UserAchievementID uint32    `json:"achievementID"`
    }
    ```