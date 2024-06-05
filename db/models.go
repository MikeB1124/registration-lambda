package db

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	LastLogin string `json:"last_login"`
}
