package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type User struct {
	Person
	Login string
}

func main() {
	u := new(User)
	// The compiler translates this u.Name to u.Person.Name
	u.Name = "Rodrigo"  // HL
	fmt.Println(u.Name) // HL
}
