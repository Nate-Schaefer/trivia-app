CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    games_played INTEGER NOT NULL DEFAULT 0,
    wins INTEGER NOT NULL DEFAULT 0,
    losses INTEGER NOT NULL DEFAULT 0,
    questions_correct INTEGER NOT NULL DEFAULT 0,
    questions_incorrect INTEGER NOT NULL DEFAULT 0,
    win_streak INTEGER NOT NULL DEFAULT 0,
    best_win_streak INTEGER NOT NULL DEFAULT 0,
    correct_streak INTEGER NOT NULL DEFAULT 0,
    best_correct_streak INTEGER NOT NULL DEFAULT 0,
    games_hosted INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL DEFAULT 'lobby' CHECK (status IN ('lobby', 'active', 'finished')),
    host_id INTEGER NOT NULL REFERENCES players(id),
    red_score INTEGER NOT NULL DEFAULT 0,
    blue_score INTEGER NOT NULL DEFAULT 0,
    winner TEXT CHECK (winner IN ('red', 'blue')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    finished_at TIMESTAMP
);

CREATE TABLE game_players (
    game_id INTEGER NOT NULL REFERENCES games(id),
    player_id INTEGER NOT NULL REFERENCES players(id),
    team TEXT CHECK (team IN ('red', 'blue')),
    PRIMARY KEY (game_id, player_id)
);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    game_id INTEGER NOT NULL REFERENCES games(id),
    category TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    question_text TEXT NOT NULL,
    answer TEXT NOT NULL,
    asked BOOLEAN NOT NULL DEFAULT FALSE,
    "order" INTEGER NOT NULL
);

CREATE TABLE rounds (
    id SERIAL PRIMARY KEY,
    game_id INTEGER NOT NULL REFERENCES games(id),
    question_id INTEGER NOT NULL REFERENCES questions(id),
    buzzer_player_id INTEGER REFERENCES players(id),
    buzzer_team TEXT CHECK (buzzer_team IN ('red', 'blue')),
    buzzer_timestamp TIMESTAMP,
    correct BOOLEAN,
    answered_by_player_id INTEGER REFERENCES players(id),
    point_awarded_to TEXT CHECK (point_awarded_to IN ('red', 'blue'))
);
