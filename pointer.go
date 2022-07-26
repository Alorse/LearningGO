package main

import "fmt"

func main() {
	/* Punteros
	1. Es uan direccion en memoria.
	2. En lugar del valor, obtenemos la direccion de este.
	3. X y Y apuntan al mismo lugar en memoria (asd123); X,Y => asd123 => 5
	4. X => asd123 => 6
	5. Y ??? => 6
	*T => *int, *string, *bool
	Valor by default => nil
	*/

	var x, y *int

	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)

}
