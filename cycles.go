package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello world %d\n", i+1)
	}

	counter := 0
	for counter < 10 {
		fmt.Println(counter + 1)
		counter++
	}

	cForever := 0
	for {
		fmt.Println(cForever)
		cForever++
	}
}
