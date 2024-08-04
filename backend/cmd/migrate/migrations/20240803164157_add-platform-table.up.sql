CREATE TABLE platforms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    imgurl VARCHAR(255),
    release_year INTEGER
);