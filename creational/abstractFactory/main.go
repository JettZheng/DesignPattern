package main

import "fmt"

func main() {
	adidasFactory, _ := getFactory("adidas")
	shoe := adidasFactory.makeShoe()
	fmt.Println(shoe)
}
