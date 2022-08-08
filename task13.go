package main

import "fmt"

func task13() {
	fmt.Println("------------Task 13----------------")
	var a, b = 4, 3
	fmt.Println(a, " ", b)
	a, b = b, a
	fmt.Println(a, " ", b)
}
