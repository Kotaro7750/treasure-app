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

func FindArticleWithComment(db *sqlx.DB, id int64) (*model.ArticleWithComment, error) {
	a := model.ArticleWithComment{}
	article := model.Article{}
	comments := []model.Comment{}
	if err := db.Get(&article, `SELECT id, title, body, user_id FROM article WHERE id = ?`, id); err != nil {
		return nil, err
	}

	rows, err := db.Queryx("SELECT id, user_id, article_id, body FROM comment WHERE article_id = ?", id)
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

	a.Article = article
	a.Comment = comments
	return &a, nil
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
	if err := db.Select(&a, `SELECT article.id, article.title, article.body, article.user_id FROM article INNE JOIN article_tag ON article.id = article_tag.article_id WHERE article_tag.tag_id = ?`); err != nil {
		return nil, err
	}
	return a, nil
}
