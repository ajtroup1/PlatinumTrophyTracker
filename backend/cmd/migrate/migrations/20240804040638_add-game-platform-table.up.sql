CREATE TABLE game_platforms (
    game_id INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    platform_id INTEGER NOT NULL REFERENCES platforms(id) ON DELETE CASCADE,
    PRIMARY KEY (game_id, platform_id)
);