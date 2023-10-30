package main

import "fmt"

func main() {
	var n, p, q, a int
	fmt.Scanln(&n)
	var levels = make(map[int]int)
	fmt.Scan(&p)

	for i := 0; i < p; i++ {
		fmt.Scan(&a)
		levels[a]++
	}

	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		fmt.Scan(&a)
		levels[a]++
	}

	if len(levels) >= n {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}
