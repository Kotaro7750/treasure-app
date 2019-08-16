package model

type Jiro struct {
	ID      int64  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Address string `db:"address" json:"address"`
	Status  string `db:"status" json:"status"`
}

type JiroDetail struct {
	ID      int64  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Address string `db:"address" json:"address"`
	Detail  string `db:"detail" json:"detail"`
	Status  string `db:"status" json:"status"`
}

type JiroNearest struct {
	Jiro  Jiro   `json:"jiro"`
	Route string `json:"route"`
	Position string `json:"position"`
}
