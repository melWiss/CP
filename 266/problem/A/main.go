package main

import "fmt"

func main() {
	var n int
	var series string
	var arr []rune
	fmt.Scanln(&n)
	fmt.Scanln(&series)
	arr = []rune(series)

	var stonesToTake int

	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			stonesToTake++
		}
	}

	fmt.Println(stonesToTake)
}
