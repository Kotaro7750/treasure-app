-- +goose Up
CREATE TABLE jiro 
(
    id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    status int(10) NOT NULL,
    ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT jiro_tag_fk_status FOREIGN KEY(status) REFERENCES status(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE article_jiro 
(
    id int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    jiro_id int(10) UNSIGNED NOT NULL,
    article_id int(10) UNSIGNED NOT NULL,
    ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    CONSTRAINT article_jiro_fk_jiro FOREIGN KEY(jiro_id) REFERENCES jiro(id),
    CONSTRAINT article_jiro_fk_article FOREIGN KEY(article_id) REFERENCES article(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
DELETE FROM article_jiro;
DROP TABLE article_jiro;

DELETE FROM jiro;
DROP TABLE jiro;

