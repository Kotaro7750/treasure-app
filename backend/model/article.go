package model

type Article struct {
	ID     int64  `db:"id" json:"id"`
	UserID *int64 `db:"user_id" json:"user_id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
}

type ArticlePost struct {
	Article Article `json:"article"`
	Tags    []int64 `json:"tags"`
	Jiro    int64   `json:"jiro"`
}

type ArticleDetail struct {
	Article Article   `json:"article"`
	Tag     []Tag     `json:"tag"`
	Comment []Comment `json:"comment"`
	Jiro    Jiro      `json:"jiro"`
}
