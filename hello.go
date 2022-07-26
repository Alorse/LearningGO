package main

import (
	"fmt"
	"strconv"
)

// var x, y, z int
var cadena string

// var flag bool
// var cadenas []string

func main() {
	cadena = "15"
	r := 23
	r = r + 1
	r_str := strconv.Itoa(r)
	cadena_int, _ := strconv.Atoi(cadena)
	var edad int

	fmt.Println("Este es un número " + r_str)
	fmt.Println(r)
	fmt.Println(cadena_int + 12)
	fmt.Printf("Mi edad es %v\n", cadena)

	fmt.Println("Coloca tu edad:")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi edad es %d\n", edad)
}

func main2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello world %d\n", i+1)
	}
}
