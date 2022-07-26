package main

import "fmt"

type User struct {
	age            int
	name, lastname string
}

func main() {
	fredo := User{age: 34, name: "Alfredo", lastname: "Ortegón"}
	cris := new(User) // Es un puntero

	cris.name = "Cristina"

	fmt.Println(fredo)
	fmt.Println(*cris)
}
