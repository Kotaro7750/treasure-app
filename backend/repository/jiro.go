package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func FindJiro(db *sqlx.DB, id int64) (*model.Jiro, error) {
	j := model.Jiro{}
	if err := db.Get(&j, `
SELECT jiro.id, jiro.name, jiro.address, detail,status.name AS status FROM jiro INNER JOIN status ON jiro.status = status.id WHERE jiro.id = ?
`, id); err != nil {
		return nil, err
	}
	return &j, nil
}
