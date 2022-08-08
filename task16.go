package main

import (
	"fmt"
)

func task16() {
	fmt.Println("------------Task 16----------------")
	//var arr = []int{9, 1, 7, 3, 5, 4, 6, 2, 8, 0}
	var arr1 = []int{9, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0, 1, 7, 3, 5, 4, 6, 2, 8, 0}
	fmt.Println(arr1)
	quickSort(arr1)
	fmt.Println(arr1)
}

func quickSort(a []int) {
	if len(a) < 2 {
		return
	}

	pivot := len(a) / 2
	left, right := 0, len(a)-1

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
	//fmt.Println(a)
	return
}
