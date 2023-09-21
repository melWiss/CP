package main

import "fmt"

func main() {
	var a1, a2, a3, a4 int
	var sequence string

	fmt.Scanln(&a1, &a2, &a3, &a4)
	fmt.Scanln(&sequence)

	var amap = make(map[rune]int, 4)

	amap['1'] = a1
	amap['2'] = a2
	amap['3'] = a3
	amap['4'] = a4

	var score int

	for _, r := range sequence {
		score += amap[r]
	}

	fmt.Println(score)

}
