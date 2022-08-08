package main

import (
	"fmt"
	"sync"
)

func task18() {
	fmt.Println("------------Task 18----------------")
	var wg sync.WaitGroup
	var counter Counter

	fmt.Println(counter.Value())
	wg.Add(5)
	go worker(&counter, 1, &wg)
	go worker(&counter, 2, &wg)
	go worker(&counter, 3, &wg)
	go worker(&counter, 4, &wg)
	go worker(&counter, 5, &wg)

	wg.Wait()
	fmt.Println(counter.Value())
}

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()

	c.num += 1
}

func (c *Counter) Value() int {
	return c.num
}

func worker(c *Counter, n int, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		c.Inc()
		fmt.Println("Worker ", n, " increases")
	}
	wg.Done()
}
