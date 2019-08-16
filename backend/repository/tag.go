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
