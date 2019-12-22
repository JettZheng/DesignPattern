package main

type normalBuilder struct {
	windowType string
	doorType string
	floor int
}

func (b *normalBuilder)setWindowType()  {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder)setDoorType()  {
	b.doorType = "Wooden Door"
}

func (b * normalBuilder)setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder)build() house {
	return house{
		doorType: b.doorType,
		windowType: b.windowType,
		floor: b.floor,
	}
}