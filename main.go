package main

import "fmt"

func PowerH(base, power int) int {

	sum := 1
	for i := 0; i < power; i++ {
		sum *= base

	}
	return sum
}

func main() {

	fmt.Println(PowerH(2, 5))
}
