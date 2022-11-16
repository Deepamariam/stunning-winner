package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	var userInputTitle string
	var toContinue string
	var toContinueLower string

	fmt.Printf("\n\t\tStone Paper Scissor!\n\n")

	for {
		userInput := getInput()
		userInputTitle = strings.Title(userInput)
		if userInputTitle == "Stone" || userInputTitle == "Scissor" || userInputTitle == "Paper" {

			fmt.Printf("\tYou chose %q.\n\n", strings.Title(userInput))
			systemInput := computerInput()
			fmt.Printf("\tComputer chose %q.\n\n", systemInput)
			compareInput(userInputTitle, systemInput)
		Continue:
			fmt.Println("Do you want to you continue playing? Yes or No!")
			fmt.Scanln(&toContinue)
			toContinueLower = strings.ToLower(toContinue)
			if toContinueLower == "yes" || toContinueLower == "y" {
				continue
			} else if toContinueLower == "no" || toContinueLower == "n" {
				fmt.Printf("\n\t\t\tThank you for playing!\n\n")
				break
			} else {
				fmt.Println("\t\t\t\nInvalid Input. Say Yes or No!")
				goto Continue
			}
		} else {
			fmt.Printf("Invalid Choice. Choose Stone, Paper, or Scissor!\n\n")
			continue
		}
	}
}

func getInput() string {

	var userInput string
	fmt.Println("Choose Stone, Paper, or Scissor!")
	fmt.Scanln(&userInput)
	return userInput

}

func computerInput() string {

	choice := []string{"Stone", "Paper", "Scissor"}
	computerInput := choice[rand.Intn(len(choice))]
	return computerInput

}

func compareInput(inputOne, inputTwo string) {

	switch inputOne {
	case "Scissor":
		if inputTwo == "Paper" {
			fmt.Println("\nYay, you won the game!")
		} else if inputTwo == "Stone" {
			fmt.Println("\nOops, you lost the game!")
		} else {
			fmt.Printf("\nIt's neither a win nor lose!\nno")
		}

	case "Paper":
		if inputTwo == "Stone" {
			fmt.Println("\nYay, you won the game!")
		} else if inputTwo == "Scissor" {
			fmt.Println("\nOops, you lost the game!")
		} else {
			fmt.Printf("\nIt's neither a win nor lose!")
		}

	case "Stone":
		if inputTwo == "Scissor" {
			fmt.Printf("\nYay, you won the game!\n")
		} else if inputTwo == "Paper" {
			fmt.Printf("\nOops, you lost the game!\n")
		} else {
			fmt.Printf("\nIt's neither a win nor lose!\n")
		}
	}
}
