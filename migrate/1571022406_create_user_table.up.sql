CREATE TABLE IF NOT EXISTS users(
    name VARCHAR(16) PRIMARY KEY,
    password TEXT NOT NULL,
    is_deleted BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL
);
