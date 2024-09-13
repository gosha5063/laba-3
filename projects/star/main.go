package main

import "fmt"

func main() {
	var s, ans string
	fmt.Scan(&s)
	ans = ""
	for i := 0; i < len(s)-1; i++ {
		ans += string(s[i]) + "*"
	}
	ans += s[len(s)-1:]
	fmt.Println(ans)
}
