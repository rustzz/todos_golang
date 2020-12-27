package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Note struct {
	ID      uint   `json:"id"`
	Hash    string `json:"hash"`
	Owner   string `json:"owner"`
	Title   string `json:"title"`
	Text    string `json:"text"`
	Checked string `json:"checked"`
	Parent  string `json:"parent"`
}

type Data struct {
	ID      uint
	Hash    string
	Title   string
	Text    string
	Checked bool
	Owner   string
	Parent  int
}
