-- +goose Up
CREATE TABLE users (
                       user_id VARCHAR(36) PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       phone VARCHAR(20),
                       address JSONB,
                       status VARCHAR(20) NOT NULL,
                       preferences JSONB,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL
);

CREATE TABLE auth (
                      auth_id VARCHAR(36) PRIMARY KEY,
                      user_id VARCHAR(36) REFERENCES users(user_id),
                      user_type VARCHAR(20) NOT NULL,
                      email VARCHAR(255) UNIQUE NOT NULL,
                      password_hash VARCHAR(255) NOT NULL,
                      created_at TIMESTAMP NOT NULL,
                      updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE auth;
DROP TABLE users;
