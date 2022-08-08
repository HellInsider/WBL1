package main

import "fmt"

func task12() {
	fmt.Println("------------Task 12----------------")
	var data = []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Print(makeSet(data))
}

func makeSet(data []string) []string {

	set := make(map[string]bool)

	for _, str := range data {
		set[str] = true
	}
	var res []string

	for key := range set {
		res = append(res, key)
	}

	return res
}
