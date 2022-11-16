package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {

	var emp Employee = Employee{

		Id:     100,
		Name:   "Mariam",
		Salary: 5000,
	}

	fmt.Printf("%#v", emp)
	fmt.Println()
	fmt.Println(emp.Id, emp.Name, emp.Salary)

}
