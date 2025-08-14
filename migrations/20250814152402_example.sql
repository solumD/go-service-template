-- +goose Up
CREATE TABLE entity (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE entity;