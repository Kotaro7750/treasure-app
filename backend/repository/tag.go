package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllTag(db *sqlx.DB) ([]model.Tag, error) {
	t := make([]model.Tag, 0)
	if err := db.Select(&t, `SELECT id, name FROM tag`); err != nil {
		return nil, err
	}
	return t, nil
}

func CreateTag(db *sqlx.Tx, t *model.Tag) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO tag (id, name) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.ID, t.Name)
}

func AppendTag(db *sqlx.Tx, articleID int64, tagID int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_tag (article_id, tag_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleID, tagID)
}

func FindTagOfArticle(db *sqlx.DB, articleID int64) ([]model.Tag, error) {
	tags := []model.Tag{}

	rows, err := db.Queryx("SELECT article_tag.id, tag.name AS name FROM tag INNER JOIN article_tag ON tag.id = article_tag.tag_id WHERE article_id = ?", articleID)
	if err != nil {
		return nil, err
	}

	var tag model.Tag
	for rows.Next() {
		err := rows.StructScan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil

}
