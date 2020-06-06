
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE companies (
   id          SERIAL PRIMARY KEY,
   name        VARCHAR(30) NOT NULL,
   description VARCHAR(255),
   active      BOOLEAN,
   website     VARCHAR(255),
   created_at  TIMESTAMP,
   updated_at  TIMESTAMP,
   deleted_at  TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS companies;