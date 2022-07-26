package main

import "fmt"

type Human struct {
	name string
}

func (this Human) hablar() string {
	return "Bla bla bla"
}
func (this Tutor) hablar() string {
	return this.Human.hablar() + "Bienvenidos"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func main() {
	tutor := Tutor{Human{"Fredo"}, Dummy{"Cris"}}

	fmt.Println(tutor.hablar())
	fmt.Println(tutor.Human.name)
}
