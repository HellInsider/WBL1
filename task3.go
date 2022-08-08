package main

import (
	"fmt"
	"os"
	"sync"
)

func task3() {
	fmt.Println("------------Task 3----------------")
	var arr = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var channel = make(chan int)
	var sum int

	for _, i := range arr {
		wg.Add(1)
		go func(n int, ch chan int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, n*n)
			ch <- n * n
		}(i, channel)
		sum += <-channel
	}
	wg.Wait()
	fmt.Println(sum)
	close(channel)

}
