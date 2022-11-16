package main

import "fmt"

func incrementor() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	var next = incrementor()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

}
