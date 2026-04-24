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

func GetGameByCode(db *pgxpool.Pool, code string) (Game, error) {
	var g Game
	err := db.QueryRow(
		context.Background(),
		"SELECT id, code, status, host_id, red_score, blue_score FROM games WHERE code = $1",
		code,
	).Scan(&g.ID, &g.Code, &g.Status, &g.HostID, &g.RedScore, &g.BlueScore)
	return g, err
}

func JoinGame(db *pgxpool.Pool, gameID, playerID int) error {
	_, err := db.Exec(
		context.Background(),
		"INSERT INTO game_players (game_id, player_id) VALUES ($1, $2)",
		gameID, playerID,
	)
	return err
}

func IsPlayerInGame(db *pgxpool.Pool, gameID, playerID int) (bool, error) {
	var count int
	err := db.QueryRow(
		context.Background(),
		"SELECT COUNT(*) FROM game_players WHERE game_id = $1 AND player_id = $2",
		gameID, playerID,
	).Scan(&count)
	return count > 0, err
}

func SelectTeam(db *pgxpool.Pool, gameID, playerID int, team string) error {
	_, err := db.Exec(
		context.Background(),
		"UPDATE game_players SET team = $1 WHERE game_id = $2 AND player_id = $3",
		team, gameID, playerID,
	)
	return err
}
