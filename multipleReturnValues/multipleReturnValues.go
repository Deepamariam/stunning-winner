package main

import "fmt"

func addSub(a, b int) (int, int) {

	var sum = a + b
	var difference = a - b
	return sum, difference

}

func main() {

	x, y := addSub(300, 400)
	fmt.Println("Sum : ", x)
	fmt.Println("Difference :", y)
}
