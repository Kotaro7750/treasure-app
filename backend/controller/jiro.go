package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Jiro struct {
	dbx *sqlx.DB
}

func NewJiro(dbx *sqlx.DB) *Jiro {
	return &Jiro{dbx: dbx}
}
func (j *Jiro) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	jiros, err := repository.AllJiro(j.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, jiros, nil
}

func (j *Jiro) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	jid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	jiroService := service.NewJiroService(j.dbx)
	jiro, err := jiroService.Show(jid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, jiro, nil
}
