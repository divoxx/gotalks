package main

import (
	"fmt"
	"testing"
)

func TestFullName(t *testing.T) {
	u := User{FirstName: "Rodrigo", LastName: "Kochenburger"}

	if u.FullName() != "Rodrigo Kochenburger" {
		t.Error("Expected name to be concatenation of FirstName and LastName")
	}
}

type User struct {
	FirstName, LastName string
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
