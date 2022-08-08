package main

import "fmt"

func task14() {
	fmt.Println("------------Task 14----------------")
	var x interface{} = 2.3
	var y interface{} = "str"

	getType(x)
	getType(y)
}

func getType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("int:", v)
	case float64:
		fmt.Println("float64:", v)
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case chan int:
		fmt.Println("chan int:", v)
	default:
		fmt.Println("unknown")
	}
}
