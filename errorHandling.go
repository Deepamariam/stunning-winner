package main

import (
	"fmt"
)

func main() {

	fmt.Println(divide(10, 0))

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide '%d' by zero", a)
	}
	return a / b, nil
}
