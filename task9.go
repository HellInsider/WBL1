package main

import (
	"fmt"
	"sync"
)

func task9() {
	fmt.Println("------------Task 9----------------")
	mass := []int{1, 3, 4, 5, 8, 6, 5, 8, 3, 11}
	var wg sync.WaitGroup

	sender := make(chan int)
	receiver := make(chan int)
	wg.Add(2)

	go makeSquare(sender, receiver, &wg)
	go printSquare(receiver, &wg)

	for _, val := range mass {
		fmt.Println(val)
		sender <- val
	}
	close(sender)

	wg.Wait()
}

func makeSquare(inputCh chan int, outputChan chan int, wg *sync.WaitGroup) {
	for range inputCh {
		val := <-inputCh
		outputChan <- val * val
	}
	close(outputChan)
	wg.Done()

}

func printSquare(inputCh chan int, wg *sync.WaitGroup) {
	for range inputCh {
		fmt.Println(<-inputCh)
	}
	wg.Done()
}
