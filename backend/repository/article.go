package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllArticle(db *sqlx.DB) ([]model.Article, error) {
	a := make([]model.Article, 0)
	if err := db.Select(&a, `SELECT id, title, body, user_id FROM article`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindArticle(db *sqlx.DB, id int64) (*model.Article, error) {
	a := model.Article{}
	if err := db.Get(&a, `
SELECT id, title, body, user_id FROM article WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func FindComment(db *sqlx.DB, articleID int64) ([]model.Comment, error) {
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

func CreateArticle(db *sqlx.Tx, a *model.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article (title, body, user_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, a.UserID)
}

func AppendTag(db *sqlx.Tx,articleID int64, tagID int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_tag (article_id, tag_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleID,tagID)
}

func UpdateArticle(db *sqlx.Tx, id int64, a *model.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article SET title = ?, body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, id)
}

func DestroyArticle(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

//tag
func TagArticle(db *sqlx.DB, tagID int64) ([]model.Article, error) {
	a := make([]model.Article, 0)
	rows, err := db.Queryx(`SELECT article.id, article.title, article.body, article.user_id FROM article INNER JOIN article_tag ON article.id = article_tag.article_id WHERE article_tag.tag_id = ?`, tagID)
	if err != nil {
		return nil, err
	}
	var article model.Article
	for rows.Next() {
		err := rows.StructScan(&article)
		if err != nil {
			return nil, err
		}
		a = append(a, article)
	}
	return a, nil
}
