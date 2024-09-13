package main

import (
	input "fmt"
	"strconv"
	"strings"
)

func main() {
	var s string
	input.Scan(&s)
	for i := 9; i > -1; i-- {
		if strings.Contains(s, strconv.Itoa(i)) {
			input.Println(i)
			break
		}
	}
}
