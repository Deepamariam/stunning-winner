package main

import "fmt"

func main() {

	a := make([]string, 3)
	a[0] = "Merry"
	a[1] = "Mia"
	a[2] = "Mariam"
	fmt.Println(a)

	a = append(a, "Matilda")
	fmt.Println(a)

	b := []string{"Hail", "Harry", "Henna", "Hia"}
	fmt.Println(b)
}
