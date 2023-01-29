package entities

type CategoryPattern struct {
	Id         int    `db:"id"`
	Pattern    string `db:"pattern"`
	CategoryId int    `db:"category_id"`
}
