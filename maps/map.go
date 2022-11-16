package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["One"] = 1
	m["Two"] = 2
	m["Three"] = 3

	fmt.Println(m)

	//To check if a key exists in a map
	_, checkMap := m["Four"]
	fmt.Println(checkMap)

	n := map[string]int{"Foo": 1, "Bar": 2, "Baz": 3}
	fmt.Println(n)

}
