package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	fruits := map[string]string{"a": "banana", "b": "mango", "c": "apple"}
	for key, value := range fruits {
		fmt.Printf("%s --> %s\n", key, value)
	}

}
