package entities

type Category struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
