package main

import "fmt"

func main() {

	var choice, number1, number2, result int

	for {
		fmt.Printf("1. Add\n 2. Subtract\n 3. Multiply\n 4. Divide\n 5. Exit\n")

		fmt.Printf("Enter your choice: \n")
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Thank you!")
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid Choice")
			continue
		}

		fmt.Println("Enter the operands: ")
		fmt.Scanln(&number1, &number2)

		switch choice {
		case 1:
			result = number1 + number2

		case 2:
			result = number1 - number2

		case 3:
			result = number1 * number2

		case 4:
			result = number1 / number2
		}

		fmt.Printf("Result = %d\n", result)

	}

}
