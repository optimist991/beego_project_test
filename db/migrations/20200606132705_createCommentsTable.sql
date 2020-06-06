-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE comments (
  id          SERIAL PRIMARY KEY,
  aim_id      INT NOT NULL,
  body        VARCHAR(255),
  active      BOOLEAN,
  created_at  TIMESTAMP,
  updated_at  TIMESTAMP,
  deleted_at  TIMESTAMP,
  FOREIGN KEY (aim_id) REFERENCES aims(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS comments;