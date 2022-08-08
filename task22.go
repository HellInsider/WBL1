package main

import (
	"fmt"
	"math/big"
)

func task22() {
	fmt.Println("------------Task 22----------------")

	//Int64 хватит для [-2^63; 2^63]
	var a, b int64

	a = 2 << 25
	b = 2 << 30

	fmt.Println("b - a = ", b-a)
	fmt.Println("b + a = ", b+a)
	fmt.Println("b / a = ", b/a)
	fmt.Println("b * a = ", b*a)

	//Если вдруг этого не достаточно, можно использовать math.big:
	a2 := big.NewInt(a)
	b2 := big.NewInt(b)

	temp := new(big.Int)

	fmt.Println("b2 - a2 = ", temp.Sub(b2, a2))
	fmt.Println("b2 + a2 = ", temp.Add(b2, a2))
	fmt.Println("b2 / a2 = ", temp.Div(b2, a2))
	fmt.Println("b2 * a2 = ", temp.Mul(b2, a2))
}
