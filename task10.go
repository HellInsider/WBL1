package main

import "fmt"

func task10() {
	fmt.Println("------------Task 10----------------")
	array := []float64{-273.0, -25.4, -27.0, -10, -9, 13.0, 19.0, 15.5, 10, 24.5, 20, -21.0, 32.5, 4, 273}
	gist := make(map[int][]float64)
	for _, value := range array {
		firstNum := int(value / 10) //Adding to "gist" by first n-1 numbers
		if value < 0 && value-float64(firstNum*10) != 0 {
			firstNum -= 1
		}
		gist[firstNum*10] = append(gist[firstNum*10], value)
	}
	fmt.Println(gist)
}
