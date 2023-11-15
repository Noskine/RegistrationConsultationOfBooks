package book

type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int8   `json:"quantity"`
}
