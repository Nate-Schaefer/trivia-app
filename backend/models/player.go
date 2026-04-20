package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Player struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PlayerCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func CreatePlayer(db *pgxpool.Pool, email, username, passwordHash string) (Player, error) {
	var p Player
	err := db.QueryRow(
		context.Background(),
		"INSERT INTO players (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id, email, username",
		email, username, passwordHash,
	).Scan(&p.ID, &p.Email, &p.Username)
	return p, err
}

func GetPlayerByEmail(db *pgxpool.Pool, email string) (Player, string, error) {
	var p Player
	var passwordHash string
	err := db.QueryRow(
		context.Background(),
		"SELECT id, email, username, password_hash FROM players WHERE email = $1",
		email,
	).Scan(&p.ID, &p.Email, &p.Username, &passwordHash)
	return p, passwordHash, err
}
