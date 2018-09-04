
-- +migrate Up
CREATE TABLE users
  (id BIGINT, name varchar(256));

-- +migrate Down
DROP TABLE users;
