CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    release_date VARCHAR(50),
    background_img VARCHAR(255),
    rating FLOAT,
    rating_count INTEGER,
    esrb_rating VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);