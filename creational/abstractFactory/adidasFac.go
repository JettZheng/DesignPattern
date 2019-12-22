package main

type adidasFac struct {
}

func (a *adidasFac) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}
