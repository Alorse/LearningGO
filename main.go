package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)

	name := "Talentpitch"
	amountContents := 1900
	fmt.Printf("%s tiene más de %d contenidos\n", name, amountContents)

	message := fmt.Sprintf("%s tiene más de %d contenidos\n", name, amountContents)
	fmt.Println(message)

	fmt.Printf("helloMessage type is %T\n", helloMessage)
	fmt.Printf("amountContents type is %T\n", amountContents)
}
