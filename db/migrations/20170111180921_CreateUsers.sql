
-- +goose Up
CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);


-- +goose Down
DROP TABLE users;
