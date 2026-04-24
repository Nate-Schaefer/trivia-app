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
