package main

import "fmt"

func task8() {
	fmt.Println("------------Task 8----------------")
	var current int64
	var bitNum, val int
	fmt.Println("Enter number, bit number and new value")
	fmt.Scan(&current, &bitNum, &val)
	newNum := changeBit(current, bitNum, val)
	fmt.Printf("%b\n", current)
	fmt.Printf("%b\n", newNum)

}

func changeBit(current int64, bitNum int, newVal int) int64 {
	if newVal == 0 {
		current = current &^ (1 << (bitNum - 1)) //and ~
	} else {
		current = current | (1 << (bitNum - 1)) // or
	}
	return current
}
