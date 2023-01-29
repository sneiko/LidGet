package requests

type GetAllFromTo struct {
	From int `form:"from"`
	To   int `form:"to"`
}
