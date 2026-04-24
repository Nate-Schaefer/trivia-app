package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nathanschaefer/trivia-app/backend/models"
	"github.com/nathanschaefer/trivia-app/backend/utils"
)

func CreateGame(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := utils.ValidateJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		game, err := models.CreateGame(db, claims.PlayerID)
		if err != nil {
			http.Error(w, "Failed to create game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(game)
	}
}

func JoinGame(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := utils.ValidateJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var body struct {
			Code string `json:"code"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if body.Code == "" {
			http.Error(w, "Code is required", http.StatusBadRequest)
			return
		}

		game, err := models.GetGameByCode(db, body.Code)
		if err != nil {
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}

		if game.Status != "lobby" {
			http.Error(w, "Game has already started", http.StatusConflict)
			return
		}

		err = models.JoinGame(db, game.ID, claims.PlayerID)
		if err != nil {
			http.Error(w, "Failed to join game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
	}
}

func SelectTeam(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := utils.ValidateJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var body struct {
			GameID int    `json:"game_id"`
			Team   string `json:"team"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if body.Team != "red" && body.Team != "blue" {
			http.Error(w, "Team must be red or blue", http.StatusBadRequest)
			return
		}

		inGame, err := models.IsPlayerInGame(db, body.GameID, claims.PlayerID)
		if err != nil || !inGame {
			http.Error(w, "You have not joined this game", http.StatusForbidden)
			return
		}

		err = models.SelectTeam(db, body.GameID, claims.PlayerID, body.Team)
		if err != nil {
			http.Error(w, "Failed to select team", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
