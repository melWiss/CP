package main

import (
	"fmt"
)

const kRated = "rated"
const kUnrated = "unrated"
const kMaybe = "maybe"

func main() {
	var n, t int
	r := false
	u := false
	fmt.Scan(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		fmt.Scan(&t)
		if t != arr[i] {
			r = true
		} else if i > 0 && arr[i] > arr[i-1] {
			u = true
		}
	}

	if r {
		fmt.Println(kRated)
	} else if u {
		fmt.Println(kUnrated)
	} else {
		fmt.Println(kMaybe)
	}
}
