package main

import "fmt"

func task19() {
	fmt.Println("------------Task 19----------------")
	data := string("абвгде日日日日1234")

	convToRunes := []rune(data)
	fmt.Println(data)

	for i := 0; i < (len(convToRunes)-1)/2; i++ {
		convToRunes[i], convToRunes[len(convToRunes)-i-1] = convToRunes[len(convToRunes)-i-1], convToRunes[i]
	}

	s := string(convToRunes)
	fmt.Println(s)
}
