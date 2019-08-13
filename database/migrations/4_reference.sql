-- +goose Up
ALTER TABLE article ADD CONSTRAINT article_fk_user FOREIGN KEY (user_id) REFERENCES user(id);