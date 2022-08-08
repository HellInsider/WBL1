package main

import (
	"fmt"
	"time"
)

func task25() {
	fmt.Println("------------Task 25----------------")
	mySleep(time.Millisecond * 2500)
}

func mySleep(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
	}
}
