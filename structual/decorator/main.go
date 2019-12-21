package main

import (
	"log"
	"os"
	"time"
	"sync"
)

type piFunc func(int) float64

func Pi(n int) float64 {
	ch := make(chan float64)

	for i := 0; i <= n; i++ {
		go func(ch chan float64, k float64) {
			ch <- k * 2
		}(ch, float64(i))
	}

	res := 0.0
	res += <-ch

	return res
}

func wrapLogger(fun piFunc, logger *log.Logger) piFunc {

	return func(n int) float64 {

		defer func(t time.Time) {
			logger.Printf("took%v", time.Since(t))
		}(time.Now())

		res := fun(n)

		return res
	}

}

func wrapCache(fun piFunc, cache *sync.Map) piFunc {
	return func(n int) float64 {

		v,ok := cache.Load(n)
		if ok {
			return v.(float64)
		}

		res := fun(n)
		cache.Store(n,res)

		return res
	}
}

func main() {
	f := wrapCache(Pi,&sync.Map{})
	g := wrapLogger(f, log.New(os.Stdout, "Test ", 1))

	g(2000)
	g(2000)
}
