package main

type user struct{}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}
func (m manager) toString() string {
	return m.user.toString() + ", manager"
}
func main() {
	var m manager

	println(m.toString())
	println(m.user.toString())
}
