package listing

//User struct containing user info
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserList struct {
	Users []User `json:"users"`
}
