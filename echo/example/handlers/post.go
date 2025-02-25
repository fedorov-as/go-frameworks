package handlers

type Post struct {
	ID        int    `json:"id" param:"id"`
	Text      string `json:"text"`
	OwnerNick string `json:"owner"`
}
