package models

import (
	"context"
	"math/rand"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Game struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Status    string `json:"status"`
	HostID    int    `json:"host_id"`
	RedScore  int    `json:"red_score"`
	BlueScore int    `json:"blue_score"`
}

func generateCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func CreateGame(db *pgxpool.Pool, hostID int) (Game, error) {
	var g Game
	code := generateCode()
	err := db.QueryRow(
		context.Background(),
		"INSERT INTO games (code, host_id) VALUES ($1, $2) RETURNING id, code, status, host_id, red_score, blue_score",
		code, hostID,
	).Scan(&g.ID, &g.Code, &g.Status, &g.HostID, &g.RedScore, &g.BlueScore)
	return g, err
}
