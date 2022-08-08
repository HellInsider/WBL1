package main

import "fmt"

func task11() {
	fmt.Println("------------Task 11----------------")
	a := []int{2, 4, 6, 8}
	b := []int{7, 9, 2, 1, 4, 13, 8}
	c := []int{5, 1, 3, 5, 9}
	d := []int{5, 1, 3, 8, 9}
	findIntersection(a, b)
	findIntersection(b, c)
	findIntersection(c, d)
}

func findIntersection(a, b []int) []int {
	intersectionMap := make(map[int]bool)

	for _, value := range a {
		intersectionMap[value] = true //pushing first set to map "elem" - "bool"
	}

	var res []int

	for _, value := range b { //comparing with second set
		if intersectionMap[value] {
			intersectionMap[value] = false //here is realization in traditional set definition
			res = append(res, value)       //set with repeating elements requires map[int]int
		}
	}

	fmt.Println(res)
	return res
}
