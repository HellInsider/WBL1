package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func task6() {
	fmt.Println("------------Task 6----------------")

	//dedicated channel for exit
	data := make(chan int)
	go first(data)
	time.Sleep(time.Millisecond * 1200)
	data <- 1

	//data channel closing
	go second(data)
	for i := 0; i < 20; i++ {
		data <- i
	}
	time.Sleep(time.Millisecond * 1200)
	close(data)

	//variable pointer for closing
	isExit := false
	go third(&isExit)
	time.Sleep(time.Millisecond * 1200)
	isExit = true

	//Context action
	ctx, cancel := context.WithCancel(context.Background())
	go fourth(ctx)
	time.Sleep(1200 * time.Millisecond)
	cancel()

	time.Sleep(time.Millisecond * 200)
}

func first(data chan int) {
	for {
		select {
		case <-data:
			fmt.Println("First died")
			return
		default:
			fmt.Println("first: ", rand.Int()%100)
			time.Sleep(time.Millisecond * 250)
		}
	}
}

func second(data chan int) {
	for range data {
		fmt.Println("second: ", <-data)
	}
	fmt.Println("Second died")
}

func third(isExit *bool) {
	for {
		if *isExit == true {
			fmt.Println("Third died")
			return
		}

		fmt.Println("third: ", rand.Int()%100)
		time.Sleep(time.Millisecond * 250)
	}
}

func fourth(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // wait for cancel
			fmt.Println("Fourth died")
			return
		default:
			fmt.Println("fourth: ", rand.Int()%100)
			time.Sleep(time.Millisecond * 250)
		}
	}
}
