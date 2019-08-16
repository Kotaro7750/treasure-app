package repository

import (
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
