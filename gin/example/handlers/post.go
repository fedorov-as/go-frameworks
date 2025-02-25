package handlers

type Post struct {
	ID        int    `json:"id" uri:"id"`
	Text      string `json:"text"`
	OwnerNick string `json:"owner"`
}
