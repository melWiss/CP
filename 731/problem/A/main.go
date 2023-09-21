package main

import (
	"fmt"
	"math"
)

func main() {
	var name string
	var score int
	var current int8 = 0
	var isit int8
	fmt.Scanln(&name)

	for _, c := range name {
		isit = int8(c - 'a')
		if math.Abs(float64(isit-current)) < 13 {
			score += int(math.Abs(float64(isit - current)))
		} else {
			score += 26 - int(math.Abs(float64(isit-current)))
		}
		current = isit
	}

	fmt.Println(score)
}
