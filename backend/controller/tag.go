package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Tag struct {
	dbx *sqlx.DB
}

func NewTag(dbx *sqlx.DB) *Tag {
	return &Tag{dbx: dbx}
}

func (t *Tag) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	tags, err := repository.AllTag(t.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, tags, nil
}

func (t *Tag) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewTagService(t.dbx)
	id, err := tagService.Create(newTag)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newTag.ID = id

	return http.StatusCreated, newTag, nil
}
