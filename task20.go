package main

import (
	"fmt"
	"strings"
)

func task20() {
	fmt.Println("------------Task 20----------------")
	data := string("wdfghjk _-- 1213, где日 ?")
	invertWords(data)
}

func invertWords(s string) {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}
