package main

import (
	"fmt"
	"strings"
)

func task26() {
	fmt.Println("------------Task 26----------------")

	data1 := string("afasdfgijkg afgg") //No
	data2 := string("abcdeABCDE")       //No
	data3 := string("abcde FGH")        //Yes

	checkOnRepeat(data1)
	checkOnRepeat(data2)
	checkOnRepeat(data3)
}

func checkOnRepeat(data string) bool {
	data = strings.ToLower(data)
	repMap := make(map[rune]bool) //Использую мап руна - лог. тип для отслеживани яповторений

	str := []rune(data)
	for _, symb := range str {
		val := repMap[symb]

		if val {
			fmt.Println("No")
			return false
		} else {
			repMap[symb] = true
		}
	}

	fmt.Println("Yes")
	return true
}
