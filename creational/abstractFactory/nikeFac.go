package main

type nikeFac struct {
}

func (a *nikeFac) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}