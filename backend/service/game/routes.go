package game

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ajtroup1/platinum-trophy-tracker/config"
	"github.com/ajtroup1/platinum-trophy-tracker/models"
	"github.com/ajtroup1/platinum-trophy-tracker/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store models.GameStore
	userStore models.UserGameStore
}

func NewHandler(store models.GameStore, userStore models.UserGameStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/games", h.handleGetAllGames).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}", h.handleGetGameByID).Methods("GET")
	router.HandleFunc("/game-search", h.handleSearchForGame).Methods("POST")
	router.HandleFunc("/add-game-db/{id:[0-9]+}", h.handleAddGameToDB).Methods("POST")
}

func (h *Handler) handleGetAllGames(w http.ResponseWriter, r *http.Request) {
	gs, err := h.store.GetAllGames()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error receiving games: %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, gs)
}

func (h *Handler) handleGetGameByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid game id: %v", err))
		return
	}

	g, err := h.store.GetGameByID(uint(id))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error receiving game with id %d: %v", int(id), err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, g)
}

func (h *Handler) handleSearchForGame(w http.ResponseWriter, r *http.Request) {
	// Extract the "val" query parameter
	queryParams := r.URL.Query()
	val := queryParams.Get("val")

	// Check if the query parameter is present
	if val == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing query parameter 'val'"))
		return
	}

	// Prepare the API request URL
	rawgKey := config.Envs.RAWGKey
	apiURL := "https://api.rawg.io/api/games"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to parse RAWG API URL: %v", err))
		return
	}

	// Set query parameters
	query := reqURL.Query()
	query.Set("key", rawgKey)
	query.Set("search", val)
	query.Set("page_size", "50")
	reqURL.RawQuery = query.Encode()

	// Make the API request
	resp, err := http.Get(reqURL.String())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to make API request: %v", err))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to read response body: %v", err))
		return
	}

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		utils.WriteError(w, resp.StatusCode, fmt.Errorf("unexpected response status: %v", resp.Status))
		return
	}

	// Parse the JSON response
	var gameResponse models.RAWGGameResponse
	if err := json.Unmarshal(body, &gameResponse); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to parse JSON response: %v", err))
		return
	}

	// Create a list of ReturnSearchGamePayload objects
	var filteredGames []models.ReturnSearchGamePayload
	for _, game := range gameResponse.Results {
		filteredGame := models.ReturnSearchGamePayload{
			ID:       game.ID,
			Name:     game.Name,
			CoverURL: game.BackgroundIMG, // or other image URL field if necessary
		}
		filteredGames = append(filteredGames, filteredGame)
	}

	// Respond with the filtered list of games
	utils.WriteJSON(w, http.StatusOK, filteredGames)
}

func (h *Handler) handleAddGameToDB(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Extract the "user" query parameter
	queryParams := r.URL.Query()
	userIDStr := queryParams.Get("user")

	// Check if the query parameter is present
	if userIDStr == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing query parameter 'user'"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error parsing user ID: %v", err))
		return
	}

	// Prepare the API request URL
	rawgKey := config.Envs.RAWGKey
	apiURL := "https://api.rawg.io/api/games/" + idStr
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to parse RAWG API URL: %v", err))
		return
	}

	// Set query parameters
	query := reqURL.Query()
	query.Set("key", rawgKey)
	reqURL.RawQuery = query.Encode()

	// Make the API request
	resp, err := http.Get(reqURL.String())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to make API request: %v", err))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to read response body: %v", err))
		return
	}

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		utils.WriteError(w, resp.StatusCode, fmt.Errorf("unexpected response status: %v", resp.Status))
		return
	}

	var game models.GameResponse
	err = json.Unmarshal(body, &game)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Read initial game information into the database
	initial_game := models.Game{
		ID:            uint32(game.ID),
		RAWGID:        uint(game.ID),
		Name:          game.Name,
		Slug:          game.Slug,
		Description:   game.Description,
		ReleaseDate:   game.Released,
		BackgroundIMG: game.BackgroundImage,
		Rating:        uint(game.Metacritic),
		Website:       game.Website,
		CreatedAt:     time.Now(),
	}
	add_game, err := h.store.AddGame(initial_game)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add game: %v", err))
		return
	}

	// Read platforms into the database
	for _, platformReq := range game.Platforms {
		platform := platformReq.Platform
		name := platform.Name

		err = h.store.AddGamePlatform(name, add_game.ID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add platform: %v", err))
			return
		}
	}

	// Read genres into the database
	for _, genreReq := range game.Genres {
		genre := genreReq.Name

		err = h.store.AddGameGenre(genre, add_game.ID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add genre: %v", err))
			return
		}
	}

	// Track game for user
	err = h.userStore.TrackGame(uint32(userID), add_game.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error tracking game: %v", err))
		return
	}

	// Read achievements into the database with pagination
	apiURL = "https://api.rawg.io/api/games/" + idStr + "/achievements"
	reqURL, err = url.Parse(apiURL)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to parse RAWG API URL: %v", err))
		return
	}

	finished := false
	for !finished {
		// Set query parameters
		query = reqURL.Query()
		query.Set("key", rawgKey)
		reqURL.RawQuery = query.Encode()

		// Make the API request
		resp, err = http.Get(reqURL.String())
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to make API request: %v", err))
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to read response body: %v", err))
			return
		}

		// Check if the response status is OK
		if resp.StatusCode != http.StatusOK {
			utils.WriteError(w, resp.StatusCode, fmt.Errorf("unexpected response status: %v", resp.Status))
			return
		}

		// Decode the JSON from the body
		var data struct {
			Results []models.Achievement `json:"results"`
			Next    string               `json:"next"`
		}
		if err := json.Unmarshal(body, &data); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error decoding json: %v", err))
			return
		}

		// Insert achievements into the database
		for _, achievement := range data.Results {
			achievement.GameID = uint(add_game.ID)
			achID, err := h.store.AddAchievement(achievement)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding achievement: %v", err))
				return
			}
			if achID == -1 {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding achievement"))
				return
			}
			err = h.store.AddUserAchievement(uint32(userID), add_game.ID, uint32(achID))
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error adding achievement: %v", err))
				return
			}
		}

		// Check if there are more pages
		if data.Next == "" {
			finished = true
		} else {
			// Update the URL for the next request
			reqURL, err = url.Parse(data.Next)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to parse next page URL: %v", err))
				return
			}
		}
	}
}
