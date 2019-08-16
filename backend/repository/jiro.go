package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllJiro(db *sqlx.DB) ([]model.Jiro, error) {
	j := make([]model.Jiro, 0)
	if err := db.Select(&j, `SELECT jiro.id, jiro.name, address, status.name AS status FROM jiro INNER JOIN status on jiro.status = status.id`); err != nil {
		return nil, err
	}
	return j, nil
}

func FindJiro(db *sqlx.DB, id int64) (*model.JiroDetail, error) {
	j := model.JiroDetail{}
	if err := db.Get(&j, `
SELECT jiro.id, jiro.name, jiro.address, detail,status.name AS status FROM jiro INNER JOIN status ON jiro.status = status.id WHERE jiro.id = ?
`, id); err != nil {
		return nil, err
	}
	return &j, nil
}

func NearestJiro(db *sqlx.DB, position string) (*model.Jiro, error) {
	j := model.Jiro{}
	if err := db.Get(&j, `
SELECT jiro.id, jiro.name, jiro.address, status.name AS status FROM jiro INNER JOIN status ON jiro.status = status.id WHERE jiro.name = ?
`, "ラーメン二郎三田本店"); err != nil {
		return nil, err
	}
	return &j, nil
}

func AppendJiroToArticle(db *sqlx.Tx, articleID int64, jiroID int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_jiro (article_id, jiro_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleID, jiroID)
}

func FindJiroOfArticle(db *sqlx.DB, articleID int64) (*model.Jiro, error) {
	j := model.Jiro{}
	rows, err := db.Queryx(`
SELECT jiro.id, jiro.name ,jiro.address,status.name AS status
FROM (article_jiro INNER JOIN jiro on article_jiro.jiro_id = jiro.id) INNER JOIN status ON status.id = jiro.status 
WHERE article_id = ?
`, articleID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.StructScan(&j)
		if err != nil {
			return nil, err
		}
	}

	return &j, nil
}

func DestroyArticleJiroIntermediate(db *sqlx.Tx, articleID int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article_jiro WHERE article_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleID)

}
