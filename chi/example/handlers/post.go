package handlers

type Post struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	OwnerNick string `json:"owner"`
}
