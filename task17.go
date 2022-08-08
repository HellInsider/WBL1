package main

import "fmt"

func task17() {
	fmt.Println("------------Task 17----------------")
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15}
	elem := 5
	elem1 := 10 //missing elem in arr

	fmt.Println("Index of searched el is: ", binarySearch(arr, elem))  //5
	fmt.Println("Index of searched el is: ", binarySearch(arr, elem1)) //here will be error
}

func binarySearch(a []int, x int) int {
	var right, left = len(a) - 1, 0
	var middle int

	for right-left > 1 {
		middle = (right + left) / 2
		if x < a[middle] {
			right = middle
		} else if x > a[middle] {
			left = middle
		} else {
			break
		}
	}

	if a[middle] != x {
		fmt.Println("Not Found")
		return -1
	}

	return middle
}
