package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var magnets = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&magnets[i])
	}
	s := 1
	for i := 1; i < n; i++ {
		// fmt.Println(magnets[i], magnets[i-1])
		if magnets[i] != magnets[i-1] {
			s++
		}
	}
	fmt.Println(s)

}
