package user

type User struct {
	id    int
	login string
	email string
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Login() string {
	return u.login
}

func New(values map[string]string) *User {
	newUser := User{}
	newUser.login = values["login"]
	newUser.email = values["email"]

	return &newUser
}
