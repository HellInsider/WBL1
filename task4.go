package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func work(ind int, osSig chan os.Signal, data chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Println("Worker ", ind, " got ", d)
	}
	fmt.Println("Worker ", ind, " stopped")
	return
}

func task4() {
	fmt.Println("------------Task 4----------------")
	var num int
	var osSig chan os.Signal
	var data chan int
	var wg sync.WaitGroup
	osSig = make(chan os.Signal, 1)
	data = make(chan int)

	signal.Notify(osSig, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("Enter workers num")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		go work(i, osSig, data, &wg)
		wg.Add(1)
	}

	for {
		select {
		case <-osSig:
			{
				close(data)
				close(osSig)
				wg.Wait()
				return
			}
		default:
			{
				data <- rand.Int() % 10
				time.Sleep(time.Millisecond * 250)
			}
		}
	}
}
