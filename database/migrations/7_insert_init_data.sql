-- +goose Up
INSERT INTO status (name) VALUES("閉店");
INSERT INTO status (name) VALUES("営業中");

INSERT INTO jiro (name,address,status,detail) VALUES("ラーメン二郎三田本店","東京都港区三田2-16-4",2,"本店だよ");
INSERT INTO jiro (name,address,status,detail) VALUES("ラーメン二郎守谷店","茨城県守谷市美園4-1-5 美園ビル1F",2,"遠いけどおいしいから食べてみて");

INSERT INTO tag (name) VALUES("家二郎");
INSERT INTO tag (name) VALUES("製麺");
-- +goose Down
