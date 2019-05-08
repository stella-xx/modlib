package lib

import "fmt"

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Read() {
	fmt.Printf("User %s is reading....", u.name)
}

type Reader interface {
	Read()
}
