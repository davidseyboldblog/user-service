package adding

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserID struct {
	ID int64 `json:"id"`
}
