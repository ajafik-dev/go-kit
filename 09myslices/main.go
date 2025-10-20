package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices School in Golang")

	var fruitList = []string{}

	fmt.Printf("The type of fruitList is: %T\n", fruitList)
	fmt.Println("The length of fruitList is:", len(fruitList))

	fruitList = append(fruitList, "Apple", "Mango", "Banana")
	fmt.Println("The length of fruitList is:", len(fruitList), fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println("The length of fruitList is:", len(fruitList), fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println("The length of fruitList is:", len(fruitList), fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println("The length of fruitList is:", len(fruitList), fruitList)

	highScores := make([]int, 4)
	highScores[0] = 100
	highScores[1] = 50
	highScores[2] = 300
	highScores[3] = 400
	// highScores[4] = 500

	highScores = append(highScores, 789, 500, 450, 700)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))


	
	

}
