package lib

import "fmt"

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Read() {
	fmt.Printf("%s is reading...", u.name)
}

func (u *User) Run() {
	fmt.Printf("%s is running...", u.name)
}

type Reader interface {
	Read()
}
