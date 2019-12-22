package singleton

import "sync"

type singleObj struct{}

var instance *singleObj
var doOnce sync.Once

func GetInstance() *singleObj {
	doOnce.Do(func() {
		instance = &singleObj{}
	})

	return instance
}
