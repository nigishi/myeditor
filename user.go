package user

type User struct {
	ID    int64
	Name  String
	Email String
}

func NewUser(id int64, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
