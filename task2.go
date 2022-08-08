package main

import (
	"fmt"
	"os"
	"sync"
)

func task2() {
	fmt.Println("------------Task 2----------------")
	var arr = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, i := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Fprintln(os.Stdout, n*n)
		}(i)
	}
	wg.Wait()
}
