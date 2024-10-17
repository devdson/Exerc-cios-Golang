package main

import "fmt"

var x int
var y string
var z bool

func main() {

	x := 42
	y := "James Bond"
	z := true

	s := fmt.Sprintln(x, y, z)

	fmt.Println(s)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

/*Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".

42
"James Bond"
true
Agora demonstre os valores nestas variáveis, com:

Uma única declaração print
Múltiplas declarações print */
