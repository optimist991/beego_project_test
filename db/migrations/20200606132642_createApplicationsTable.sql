-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE applications (
  id          SERIAL PRIMARY KEY,
  aim_id      INT NOT NULL ,
  uniq_id     VARCHAR(255) unique,
  filename    VARCHAR(255),
  url         VARCHAR(255),
  description VARCHAR(255),
  active      BOOLEAN,
  created_at  TIMESTAMP,
  updated_at  TIMESTAMP,
  deleted_at  TIMESTAMP,
  FOREIGN KEY (aim_id) REFERENCES aims(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS applications;