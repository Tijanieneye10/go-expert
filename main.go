package main

import (
	"fmt"
	"go-expert/colorers"
	"html/template"
	"os"
	"strings"
	"unicode"
)

type User struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	s1 := "John Doe"
	s2 := strings.Clone(s1)

	s3 := strings.TrimSpace(s2)

	s4 := unicode.IsLetter(rune(s3[0]))

	fmt.Println(s4)

	templates := `
		Hello, {{ .Firstname }}
		My age is {{ .Age}}
	`

	tmp, err := template.New("my_template").Parse(templates)

	if err != nil {
		fmt.Println(err)
	}

	err = tmp.Execute(os.Stdout, User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       42,
	})
	if err != nil {
		return
	}

	fmt.Println(colorers.Text("Hello John Doe", colorers.Green, colorers.Red))

	value := fmt.Sprintf("%#v", s3)
	fmt.Println(value)
}
