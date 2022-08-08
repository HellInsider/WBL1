package main

import (
	"fmt"
	"sync"
)

func task7() {
	fmt.Println("------------Task 7----------------")
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(m, i, &wg, &mutex)
	}

	wg.Wait()
	fmt.Println(m)
}

func writeToMap(m map[int]int, i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Println("Goroutine num ", i, " writes")
	m[i] = i
	mutex.Unlock()
	wg.Done()
}
