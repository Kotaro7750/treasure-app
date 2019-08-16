package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateComment(db *sqlx.Tx, c *model.Comment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO comment (body, user_id,article_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.Body, c.UserID, c.ArticleID)
}

func FindCommentOfArticle(db *sqlx.DB, articleID int64) ([]model.Comment, error) {
	comments := []model.Comment{}

	rows, err := db.Queryx("SELECT id, user_id, article_id, body FROM comment WHERE article_id = ?", articleID)
	if err != nil {
		return nil, err
	}

	var comment model.Comment
	for rows.Next() {
		err := rows.StructScan(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil

}

func DestroyArticleComment(db *sqlx.Tx, articleID int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM comment WHERE article_id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleID)

}
