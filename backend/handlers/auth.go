package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"

	"github.com/nathanschaefer/trivia-app/backend/models"
	"github.com/nathanschaefer/trivia-app/backend/utils"
)

func Register(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds models.PlayerCredentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if creds.Email == "" || creds.Password == "" || creds.Username == "" {
			http.Error(w, "Email, password, and username are required", http.StatusBadRequest)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		player, err := models.CreatePlayer(db, creds.Email, creds.Username, string(hash))
		if err != nil {
			log.Printf("Failed to create player: %v", err)
			errMsg := err.Error()
			if strings.Contains(errMsg, "players_email_key") {
				http.Error(w, "Email is already registered", http.StatusConflict)
			} else if strings.Contains(errMsg, "players_username_key") {
				http.Error(w, "Username is already taken", http.StatusConflict)
			} else {
				http.Error(w, "Failed to create player", http.StatusInternalServerError)
			}
			return
		}

		token, err := utils.GenerateJWT(player.ID, player.Username)
		if err != nil {
			log.Printf("Failed to generate JWT: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func Login(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds models.PlayerCredentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if creds.Email == "" || creds.Password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		player, hash, err := models.GetPlayerByEmail(db, creds.Email)
		if err != nil {
			http.Error(w, "No account found with that email", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(creds.Password)); err != nil {
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}

		token, err := utils.GenerateJWT(player.ID, player.Username)
		if err != nil {
			log.Printf("Failed to generate JWT: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
