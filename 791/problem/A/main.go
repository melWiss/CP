package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var years = 0
	for a <= b {
		years++
		a *= 3
		b *= 2
	}
	fmt.Println(years)
}
