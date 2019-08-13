-- +goose Up
ALTER TABLE article add column user_id int
(10) UNSIGNED NOT NULL REFERENCES user
(id);