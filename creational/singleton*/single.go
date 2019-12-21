package singleton

import (
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var instance2 *single

func getInstance2() *single {
	if instance2 == nil {
		lock.Lock()
		defer lock.Unlock()
		instance2 = &single{}
	}

	return instance2

}
