-- +goose Up
CREATE TABLE jiro_detail 
(
    jiro_id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    body TEXT,
    ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (jiro_id),
    CONSTRAINT jiro_detail_fk_status FOREIGN KEY(jiro_id) REFERENCES jiro(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DELETE FROM jiro_detail;
DROP TABLE jiro_detail;
