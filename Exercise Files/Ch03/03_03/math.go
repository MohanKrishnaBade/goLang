package main

import (
	"fmt"
	"math/big"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	var int1 int = 23
	var float1 float64 = 23.9

	fmt.Println("add:", (float64(int1) + float1))

	f1, f2, f3 := 23.5, 34.6, 45.7

	floatsum := f1 + f2 + f3

	fmt.Println(floatsum)

	var b1, b2, b3, bigsum big.Float

	b1.SetFloat64(34.6)
	b2.SetFloat64(45.9)
	b3.SetFloat64(67.8)

	bigsum.Add(&b1, &b2).Add(&bigsum, &b3)

	fmt.Printf("float add: %.10g\n", &bigsum)

	somethjinvd := 45.7
	fmt.Printf("math package: %.2f\n", (somethjinvd * math.Pi))

}
