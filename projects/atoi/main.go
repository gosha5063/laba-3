package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	var queue []int
	for num > 0 {
		a := num % 10
		queue = append(queue, a*a)
		num = num / 10
	}
	for i := 0; i < len(queue); i++ {
		fmt.Print(queue[len(queue)-i-1])
	}
}
