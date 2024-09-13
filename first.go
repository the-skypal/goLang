package main

import "fmt"

type Product struct {
	Quantity int
	Price    float64
	InStock  bool
}

func main() {
	products := make(map[string]Product)

	products["Laptop"] = Product{Quantity: 10, Price: 29999.90, InStock: true}
	products["Mouse"] = Product{Quantity: 20, Price: 99.90, InStock: true}
	products["Keyboard"] = Product{Quantity: 30, Price: 199.90, InStock: false}

	for name, product := range products {
		fmt.Printf("Product: %s\n", name)
		fmt.Printf("Quantity: %d\n", product.Quantity)
		fmt.Printf("Price: $%.2f\n", product.Price)
		fmt.Printf("In Stock: %t\n\n", product.InStock)
	}

	productName := "Laptop"
	if product, exists := products[productName]; exists {
		fmt.Printf("%s is available with quantity %d\n", productName, product.Quantity)
	} else {
		fmt.Printf("%s is not available\n", productName)
	}

	fmt.Println("DONE WITH BUILD")
}
