package main

import (
	"log"
	"strconv"
)

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.Acquire()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.Acquire()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	pool.Release(conn1)
	pool.Release(conn2)
}
