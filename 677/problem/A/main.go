package main

import "fmt"

func main() {
	var n int
	var h int
	fmt.Scanln(&n, &h)
	var a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var sum = 0

	for i := 0; i < n; i++ {
		if a[i] <= h {
			sum += 1
		} else {
			sum += 2
		}
	}
	fmt.Println(sum)
}
