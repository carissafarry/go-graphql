package post

// type Post struct {
// 	ID          string `json:"id"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	User        *User  `json:"user"`
// }

type Post struct {
	ID          string
	Title       string
	Description string
	UserID      string
}