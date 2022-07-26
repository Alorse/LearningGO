package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	slice = append(slice, 2)
	fmt.Println(slice)

	array := []int{1, 2, 3, 4}
	copia := make([]int, len(array))
	/*
		Copia el mini de elementos
	*/
	copy(copia, array)

	fmt.Println(copia)
}
