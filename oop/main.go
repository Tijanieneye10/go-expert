package oop

type User struct {
	Firstname string
	Lastname  string
	Age       int
}

func (u User) FullName() string {
	return u.Firstname + " " + u.Lastname
}
