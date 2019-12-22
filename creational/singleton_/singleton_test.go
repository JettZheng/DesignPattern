package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Fatal("failed")
	}
}

func TestGetInstance2(t *testing.T) {
	var res = make(chan *single, 100)

	for i := 0; i < 100; i++ {
		go func() {
			t := getInstance2()
			res <- t
		}()
	}

	for index := 0; index < len(res); index++ {
		var a = <-res
		var b = <-res
		if a != b {
			t.Fatal("eror")
		}
	}
}
