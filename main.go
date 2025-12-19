package main

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Age       int
	Address
}

type Address struct {
	Street string
	City   string
	State  string
}

func (u *User) GetUserStreet() string {
	return fmt.Sprintf("%v", u.Street)
}

func main() {
	user := &User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
		Address: Address{
			Street: "My Street is here",
			City:   "Tokyo",
			State:  "Paris",
		},
	}

	fmt.Println(user.GetUserStreet())

}
