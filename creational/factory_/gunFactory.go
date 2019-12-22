package gunFactory

import (
	"errors"
	"fmt"
)

func Create(name string) (iGun, error) {
	switch name {
	case "AK":
		return newAk47(), nil
	case "Mave":
		return newMave(), nil
	default:
		return nil, errors.New(fmt.Sprintf("Wrong type"))
	}
}
