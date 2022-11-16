package main

import "fmt"

func main() {

	sum := plus(10, 20, 30)
	fmt.Println("Sum = ", sum)

}

func plus(a, b, c int) int {

	return a + b + c

}
