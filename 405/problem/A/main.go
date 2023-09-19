package main

import (
	"fmt"
	"sort"
)

func batata() (int, int) {
	return 5, 7
}

func main() {
	var n int
	fmt.Scanln(&n)

	var arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr[:])

	for _, i := range arr {
		fmt.Println(i)
	}

}
