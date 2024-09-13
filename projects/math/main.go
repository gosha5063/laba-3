package main

import (
	"fmt"
	"math"
)

func T(w float64) float64 {
	return 6 / w
}
func W(k, m float64) float64 {
	return math.Sqrt(k / m)
}
func M(p, v float64) float64 {
	return p * v
}
func main() {
	var k, p, v float64
	fmt.Scan(&k, &p, &v)
	fmt.Println(T(W(k, M(p, v))))
}
