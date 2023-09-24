package main

import (
	"fmt"
)

func main() {
	var n, h, a, score int
	fmt.Scanln(&n)
	var home = make(map[int]int)
	var away = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scanln(&h, &a)
		home[h]++
		away[a]++
	}

	// fmt.Println("___________________________________")
	// fmt.Println(home)
	// fmt.Println(away)
	// fmt.Println("___________________________________")
	for k, v := range home {
		score += v * away[k]
	}

	fmt.Println(score)

}
