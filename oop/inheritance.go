package main

import "fmt"

type UserDetail struct {
	Firstname string
	Lastname  string
	Age       int
	Address
	SemiAddress
}

type Address struct {
	Street string
	City   string
	State  string
}

type SemiAddress struct {
	Street string
}

func (u *UserDetail) GetUserStreet() string {
	return fmt.Sprintf("%v", u.SemiAddress.Street)
}

func main() {
	user := &UserDetail{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
		Address: Address{
			Street: "My Street is here",
			City:   "Tokyo",
			State:  "Paris",
		},
		SemiAddress: SemiAddress{
			Street: "second street",
		},
	}

	fmt.Println(user.GetUserStreet())

}
