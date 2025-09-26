package domain


type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner" db:"is_owner"`
}