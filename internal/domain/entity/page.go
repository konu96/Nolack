package entity

type Page struct {
	ID string `json:"id"`
}

func NewPage(ID string) Page {
	return Page{
		ID: ID,
	}
}
