package main

import (
	"errors"
	"fmt"
	"sync"
)

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, errors.New("can not create 0 ppool")
	}

	var activePoolObjects []iPoolObject

	pool := &pool{
		idle:     poolObjects,
		active:   activePoolObjects,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}

	return pool, nil
}

func (p *pool) Acquire() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()

	if len(p.idle) == 0 {
		return nil, errors.New("all busy")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]

	p.active = append(p.active, obj)

	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())

	return obj,nil
}

func (p *pool) Release(target iPoolObject) error {
	p.mulock.Lock()
    defer p.mulock.Unlock()
    err := p.remove(target)
    if err != nil {
        return err
    }
    p.idle = append(p.idle, target)
    fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
    return nil
}

func (p *pool) remove(target iPoolObject) error {
    currentActiveLength := len(p.active)
    for i, obj := range p.active {
        if obj.getID() == target.getID() {
            p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
            p.active = p.active[:currentActiveLength-1]
            return nil
        }
    }
    return fmt.Errorf("Targe pool object doesn't belong to the pool")
}