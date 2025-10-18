package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome time study")

	presentTime := time.Now()

	fmt.Println("Present Time ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2025,time.October, 18,10,12,23,0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
