package entity

type Product struct {
	Id    string `json:"-"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Size  int    `json:"size"`
	Count int    `json:"count"`
}
