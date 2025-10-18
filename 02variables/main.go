package main

import "fmt"

func main() {
	fmt.Println("Variables")
	var username string = "ajafik"
	fmt.Println(username)
	fmt.Printf("Variable is of Type: %T -- value %v\n", username, username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of Type: %T -- value %v\n", isLoggedIn, isLoggedIn)


	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of Type: %T \n", smallVal)

		var smallFloat float64 = 255.4557677676
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of Type: %T \n", smallFloat)

}
