package main

import (
	"fmt"
	//"strings"
)

func main() {
	var palabra string
	fmt.Scanln(&palabra)

	for i := 0; i < len(palabra); i++ {
		fmt.Println(string(palabra[i]))
	}

	fmt.Println("El largo de la palabra es: ", len(palabra))
}
