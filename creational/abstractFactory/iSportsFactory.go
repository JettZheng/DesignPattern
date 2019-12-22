package main

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
}

func getFactory(name string) (iSportsFactory, error) {
	switch name {
	case "adidas":
		return &adidasFac{}, nil
	case "nike":
		return &nikeFac{}, nil
	default:
		return nil, fmt.Errorf("wrong type")
	}
}
