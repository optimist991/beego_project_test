
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users_aims (
    id          SERIAL PRIMARY KEY,
    user_id     INT NOT NULL,
    aim_id      INT NOT NULL,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (aim_id) REFERENCES aims(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS users_aims;