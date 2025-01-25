package domain

type User struct {
	Username string
	Active   bool
}

func NewUser(username string) *User {
	return &User{
		Username: username,
		Active:   true,
	}
}
