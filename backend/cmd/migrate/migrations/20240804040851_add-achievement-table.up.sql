CREATE TABLE achievements (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    imgurl VARCHAR(255),
    percent FLOAT NOT NULL
);  