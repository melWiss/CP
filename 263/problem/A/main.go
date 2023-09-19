package main

import (
	"fmt"
	"math"
)

const n = 5

func main() {
	var m [n][n]int
	var x = 0
	var y = 99
	for i := 0; i < n; i++ {
		fmt.Scanln(&m[i][0], &m[i][1], &m[i][2], &m[i][3], &m[i][4])
		if y == 99 {
			x = i
			for j := 0; j < n; j++ {
				if m[i][j] == 1 {
					y = j
				}
			}
		}
	}

	x = int(math.Abs(float64(x - 2)))

	y = int(math.Abs(float64(y - 2)))

	fmt.Println(x + y)

}
