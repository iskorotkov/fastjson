package example

import "time"

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Links     []Link    `json:"links"`
	Friends   []Friend  `json:"friends"`
}

type Link struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type Friend struct {
	ID       int       `json:"id"`
	Since    time.Time `json:"since"`
	Nickname string    `json:"nickname"`
}
