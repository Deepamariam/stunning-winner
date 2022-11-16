package main

import "fmt"

func main() {

	var arrayA [5]int
	fmt.Println("Array is: ", arrayA) //[0, 0, 0, 0, 0]
	arrayA[0] = 1
	fmt.Println("Array A is: ", arrayA) //[0, 0, 0, 0, 0]

	arrayB := [5]int{100, 200, 300, 400, 500}
	fmt.Println("Array B is: ", arrayB)

}
