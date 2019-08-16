package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Article struct {
	db *sqlx.DB
}

func NewArticleService(db *sqlx.DB) *Article {
	return &Article{db}
}

func (a *Article) Show(articleID int64) (*model.ArticleDetail, error) {
	article, err := repository.FindArticle(a.db, articleID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find article")
	}

	comment, err := repository.FindComment(a.db, articleID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find comment")
	}

	articleDetail := model.ArticleDetail{}
	articleDetail.Article = *article
	articleDetail.Comment = comment
	//tagはまだ

	return &articleDetail, nil
}

func (a *Article) Update(id int64, newArticle *model.Article) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateArticle(tx, id, newArticle)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article update transaction")
	}
	return nil
}

func (a *Article) Destroy(id int64) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyArticle(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article delete transaction")
	}
	return nil
}

func (a *Article) Create(newArticle *model.ArticleCreate) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateArticle(tx, &newArticle.Article)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed article insert transaction")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		for _, tagID := range newArticle.Tags {
			_, err := repository.AppendTag(tx, createdId, tagID)
			if err != nil {
				return err
			}

		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, errors.Wrap(err, "failed article_tag insert transaction")
	}
	return createdId, nil
}

func (a *Article) IndexByTag(tagID int64) ([]model.Article, error) {
	return nil, nil
}
