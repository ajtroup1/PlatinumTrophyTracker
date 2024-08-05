CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    rawg_id INT NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    release_date VARCHAR(50),
    background_img VARCHAR(255),
    rating FLOAT,
    website VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);