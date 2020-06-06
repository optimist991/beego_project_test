
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE aims (
  id          SERIAL PRIMARY KEY,
  parent_id   INT,
  type        VARCHAR(55),
  title       VARCHAR(255),
  description VARCHAR(255),
  created_at  TIMESTAMP,
  updated_at  TIMESTAMP,
  deleted_at  TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS aims;