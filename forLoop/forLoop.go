package main

import "fmt"

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("**************************")

	for i = 0; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i)
		}

	}

}
