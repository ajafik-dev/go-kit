package main // This is the main package declaration

import (
	"fmt"
)

func main() {
	/**
	* This is a a mullti line comment and the code below prints "Hello World"
	 */
	var surname string = "Ajayi"
	othernames := "Oluwafikayo"

	var a,b,c int = 1,2,3
	fmt.Println(a,b,c)

	const PI = 3.142
	fmt.Println(PI)

	fmt.Println(surname, othernames)
	fmt.Println("Hello World")
}
