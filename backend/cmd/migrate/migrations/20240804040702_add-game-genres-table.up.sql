CREATE TABLE game_genres (
    game_id INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    genre VARCHAR(255) NOT NULL,
    PRIMARY KEY (game_id, genre)
);