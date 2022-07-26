package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3}
	var obj [2][2]int

	fmt.Println(obj)
	fmt.Println(len(array))

	array[3] = 4

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
