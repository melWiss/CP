package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	var a = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	var m, x, y int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&x, &y)
		if x < n {
			a[x+1] += a[x] - y
		}
		if x > 1 {
			a[x-1] += y - 1
		}
		a[x] = 0

	}
	for i := 1; i <= n; i++ {
		fmt.Println(a[i])
	}

}
