package main

import "fmt"

func main() {

	var num int = 5

	fmt.Println(num)
	fmt.Println(&num)

	var name string = "Milda"
	var memoryAddress *string = &name
	fmt.Println(memoryAddress)

}
