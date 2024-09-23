package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	culculateFib(num)
}

func culculateFib(num int) {
	firstFib := num
	a, b := 0, 1

	for b < firstFib {
		a, b = b, a+b
	}

	for i := 0; i < 10; i++ {
		fmt.Println(b)
		a, b = b, a+b
	}

}
