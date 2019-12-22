package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")

	director := newDirector(normalBuilder)

	normalHouse := director.build()

	fmt.Printf("Normal House Door Type:%v\n", normalHouse)

	iglooBuilder := getBuilder("igloo")

	director.setBuilder(iglooBuilder)

	iglooHouse := director.build()

	fmt.Printf("igloo house:%v\n",iglooHouse)
}
