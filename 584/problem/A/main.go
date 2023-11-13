package main

import (
	"fmt"
)

func main() {
	var n, t int
	var score string

	fmt.Scan(&n, &t)
	if n == 1 {
		if t == 10 {
			fmt.Println(-1)
		} else {
			fmt.Println(t)
		}
	} else if t == 10 {
		score = "1"
		for i := 0; i < n-1; i++ {
			score += "0"
		}
		fmt.Println(score)
	} else {
		for i := 0; i < n; i++ {
			score += fmt.Sprint(t)
		}
		fmt.Println(score)
	}

}
