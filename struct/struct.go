package main

import "fmt"

type Player struct {
	name  string
	age   int
	city  string
	state string
}

func main() {

	player1 := Player{

		name:  "Sam",
		age:   25,
		city:  "Hisar",
		state: "Haryana",
	}

	player2 := Player{

		name:  "Jam",
		age:   30,
		city:  "Ludhiana",
		state: "Punjab",
	}

	fmt.Println("Player 1: ", player1)
	fmt.Println("Player 2:", player2)

	fmt.Println(player1.name)
	fmt.Println(player2.name)

}
