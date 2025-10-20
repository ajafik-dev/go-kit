package main

import "fmt"

func main() {
	fmt.Println("Welcome to array study in Golang")

	var bookList [5]string

	bookList[0] = "The Alchemist"
	bookList[1] = "The Power of Now"
	bookList[2] = "The Power of Subconscious Mind"
	bookList[3] = "The Power of Positive Thinking"
	bookList[4] = "The Power of Habit"

	fmt.Println("Book List: ", bookList, len(bookList))

}
