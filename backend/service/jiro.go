package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Jiro struct {
	db *sqlx.DB
}

func NewJiroService(db *sqlx.DB) *Jiro {
	return &Jiro{db}
}

func (j *Jiro) Show(jiroID int64) (*model.DetailedJiro, error) {
	jiro, err := repository.FindJiro(j.db, jiroID)
	if err != nil {
		return nil, errors.Wrap(err, "failed find jiro")
	}

	return jiro, nil
}
