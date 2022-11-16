package main

import (
	"fmt"
	"time"
)

func main() {

	whatAmI := func(i interface{}) {

		switch t := i.(type) {

		case bool:
			fmt.Println("It is a boolean value.")
		case int:
			fmt.Println("It is an integer value.")
		default:
			fmt.Printf("Don't know type %T.\n", t)
		}
	}

	whatAmI(true)
	whatAmI(10)
	whatAmI("*")

	t := time.Now()

	switch {

	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	case t.Hour() > 12:
		fmt.Println("It's after noon.")
	case t.Hour() == 12:
		fmt.Println("It's noon.")
	default:
		fmt.Println("Invalid!")

	}

}
