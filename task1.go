package main

import "fmt"

type Human struct {
	weight, height int
}

type Action struct {
	human Human
}

func (h Human) getWeight() int {
	return h.weight
}

func (h Human) getHeight() int {
	return h.height
}

func task1() {
	fmt.Println("------------Task 1----------------")
	h := Action{human: Human{weight: 12, height: 34}}
	print(h.human.getWeight())
}
