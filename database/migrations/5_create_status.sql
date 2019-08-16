-- +goose Up
CREATE TABLE status
(
    id int(10) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4; 
-- +goose Down
DELETE FROM status;
DROP TABLE status;