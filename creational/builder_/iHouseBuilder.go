package main

type iHouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	build() house
}

func getBuilder(builderType string) iHouseBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	default:
		return nil
	}
}
