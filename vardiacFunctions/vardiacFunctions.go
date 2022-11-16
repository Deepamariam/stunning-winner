package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sums(nums...))

}

func sums(nums ...int) int {

	total := 0
	for _, num := range nums {
		total = total + num
	}

	return total
}
