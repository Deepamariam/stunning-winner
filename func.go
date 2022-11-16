package main

import (
	"fmt"
	"log"
)

func main() {

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
}

func logOperation(oper func(int, int) int, n1, n2 int) {
	log.Println("Operation started")
	result := oper(n1, n2)
	fmt.Println(result)
	log.Println("Operation completed")
}

func logAdd(n1, n2 int) {
	log.Println("Operation started")
	add(n1, n2)
	log.Println("Operation completed")
	fmt.Println(add)
}

func logSubtract(n1, n2 int) {
	log.Println("Operation started")
	subtract(n1, n2)
}

func add(n1, n2 int) int {
	result := n1 + n2
	// fmt.Println("Result :", result)
	return result
}

func subtract(n1, n2 int) int {
	result := n1 - n2
	// fmt.Println("Result :", result)
	return result
}
