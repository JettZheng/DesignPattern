package gunFactory

type ak47 struct {
	gun
}

func newAk47() iGun{
	return &ak47{
		gun: gun{
			name:"AK47",
			power:10,
		},
	}
}

