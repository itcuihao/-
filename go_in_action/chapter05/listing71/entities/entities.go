package entities

type user struct {
	Name  string
	email string
}

type Admin struct {
	user
	Rights int
}

func (u *user) SetEmail(e string) {
	u.email = e
}
