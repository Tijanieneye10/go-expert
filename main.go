package main

import (
	"fmt"
)

type User struct {
	Firstname string
	Lastname  string
	Age       int
}

type Payment interface {
	Initialize()
	Verify()
}

type Flutterwave struct {
	Amount    float64
	Reference string
}

type Paystack struct {
	Amount    float64
	Reference string
}

func (f *Flutterwave) Initialize() {
	fmt.Println("initializing flutterwave")
}

func (f *Flutterwave) Verify() {
	fmt.Println("verifying flutterwave")
}

func (f *Flutterwave) String() string {
	return fmt.Sprintf("%v", "flutterwave is here")
}

func (p *Paystack) Initialize() {
	fmt.Println("initializing paystack")
}

func (p *Paystack) Verify() {
	fmt.Println("verifying paystack")
}

func MakePayment(p Payment) {
	p.Initialize()
	p.Verify()
}

func (u User) FullName() string {
	return u.Firstname + " " + u.Lastname
}

func main() {
	f := Flutterwave{
		Amount:    10,
		Reference: "Salary-3282893892",
	}

	p := Paystack{
		Amount:    30,
		Reference: "Salary-32828924232",
	}

	fmt.Println(&f)

	MakePayment(&p)
	MakePayment(&f)
}
