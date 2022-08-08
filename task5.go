package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func read(data <-chan int, wg *sync.WaitGroup) {
	for {
		val, opened := <-data
		if !opened {
			break
		}
		fmt.Println(val)
	}
	wg.Done()

}

func write(data chan int, osSig <-chan int) {
	for {
		select {
		case <-osSig:
			return
		default:
			data <- rand.Int() % 10
			time.Sleep(time.Millisecond * 250)
		}
	}
}

func task5() { //Here I use osSig just to make trigger to stop goroutines
	fmt.Println("------------Task 5----------------")
	var seconds time.Duration
	var osSig chan int
	var data chan int
	var wg sync.WaitGroup
	osSig = make(chan int)
	data = make(chan int)
	wg.Add(1)

	fmt.Println("Enter seconds num")
	fmt.Scan(&seconds)

	go read(data, &wg)
	go write(data, osSig)

	time.Sleep(seconds * time.Second)
	osSig <- 1
	close(data)
}
