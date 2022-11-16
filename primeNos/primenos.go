package main

import "fmt"

func main() {

	var flag int

	for i := 3; i <= 100; i++ {
		flag = 0
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				flag = 1
			}
		}
		if flag == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
