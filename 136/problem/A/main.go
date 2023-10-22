package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		b[a[i]-1] = i + 1
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", b[i])
	}
}
