package main

import (
	"fmt"
)

func main() {
	var n int
	var min, max int
	fmt.Scanln(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Printf("%d %d\n", (a[1] - a[0]), (a[n-1] - a[0]))
	for i := 1; i < n-1; i++ {
		min1 := a[i] - a[i-1]
		min2 := a[i+1] - a[i]
		if min1 < min2 {
			min = min1
		} else {
			min = min2
		}
		max1 := a[n-1] - a[i]
		max2 := a[i] - a[0]
		if max1 > max2 {
			max = max1
		} else {
			max = max2
		}
		fmt.Printf("%d %d\n", min, max)
	}
	fmt.Printf("%d %d\n", (a[n-1] - a[n-2]), (a[n-1] - a[0]))
}
