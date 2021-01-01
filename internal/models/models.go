package models

// User : ...
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// Note : ...
type Note struct {
	ID      uint   `json:"id"`
	Hash    string `json:"hash"`
	Owner   string `json:"owner"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
	Parent  int    `json:"parent"`
}
