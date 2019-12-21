package main
//director decide to get which builder and start build

type director struct {
	builder iHouseBuilder
}

func newDirector(b iHouseBuilder) *director  {
	return &director{
		builder : b ,
	}
}

func (d *director)setBuilder(b iHouseBuilder)  {
	d.builder = b
}

func (d *director)build() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()

	return d.builder.build()
}