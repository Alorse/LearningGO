package main

import (
	"fmt"
	"strings"
)

func main() {
	// var array []int

	// array = []int{1, 23}

	// if array == nil {
	// 	fmt.Println("Esta vacio")
	// } else {
	// 	fmt.Println(array)
	// }

	// array2 := [4]int{1, 2, 3, 4}
	// slice := array2[1:]
	// fmt.Println(slice)

	// rslice := []string{"Tití", "me", "preguntó"}

	// for _, value := range rslice {
	// 	fmt.Println(value)
	// }

	// for i, value := range rslice {
	// 	fmt.Println(i, value)
	// }

	var palabra string
	fmt.Scan(&palabra)

	isPalindromo(palabra)
}

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
