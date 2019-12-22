package gunFactory

type mave struct {
	gun
}

func newMave() iGun {
	return &mave{
		gun: gun{
			name:  "Mave",
			power: 4,
		},
	}
}
