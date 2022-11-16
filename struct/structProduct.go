package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func printProduct(product Product) {

	fmt.Printf("Id: %d, Name: %q, Cost: %v, Category: %q", product.Id, product.Name, product.Cost, product.Category)

}

func applyDiscount(product *Product, discountPercent float32) {

	discount := (discountPercent / 100) * product.Cost
	product.Cost = product.Cost - discount

}
func main() {

	pen := Product{

		Id:       1007,
		Name:     "Pen",
		Cost:     100,
		Category: "Stationary",
	}

	printProduct(pen)
	applyDiscount(&pen, 10)
	printProduct(pen)

}
