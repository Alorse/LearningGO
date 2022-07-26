package main

import "fmt"

type User struct {
	age            int
	name, lastname string
}

func (this User) full_name() string {
	return this.name + " " + this.lastname
}

func (this *User) set_name(n string) {
	this.name = n
}

func main() {
	fredo := User{age: 34, name: "Alfredo", lastname: "Ortegón"}
	cris := new(User) // Es un puntero

	cris.name = "Cristina"
	cris.lastname = "Hernandez"
	cris.set_name("Angela")

	fmt.Println(fredo)
	fmt.Println(cris.full_name())
}
