package main

import "fmt"

func main() {
	var n, b, d, s, a, j int
	fmt.Scanln(&n, &b, &d)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a <= b {
			s += a
		}

		if s > d {
			s = 0
			j++
		}
	}

	fmt.Println(j)
}
