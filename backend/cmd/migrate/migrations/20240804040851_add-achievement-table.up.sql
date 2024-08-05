CREATE TABLE achievements (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    imgurl VARCHAR(255),
    percent VARCHAR(40) NOT NULL,
    game_id INT REFERENCES games(id) ON DELETE CASCADE
);
