CREATE TABLE screenshots (
    id SERIAL PRIMARY KEY,
    game_id INTEGER NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    imgurl VARCHAR(255) NOT NULL
);