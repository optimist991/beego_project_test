
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
   id          SERIAL PRIMARY KEY,
   company_id  INT NOT NULL,
   first_name  VARCHAR(55),
   last_name   VARCHAR(55),
   location    VARCHAR(255),
   email       VARCHAR(255) unique,
   active      BOOLEAN,
   created_at  TIMESTAMP,
   updated_at  TIMESTAMP,
   deleted_at  TIMESTAMP,
   FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
--
DROP TABLE IF EXISTS users;