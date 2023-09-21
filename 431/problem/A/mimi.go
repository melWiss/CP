package main

import "fmt"

func main() {
	var arr = make([]int, 4)
	var sequence string

	fmt.Scanln(&arr[0], &arr[1], &arr[2], &arr[3])
	fmt.Scanln(&sequence)

	var score int

	for _, c := range sequence {
		score += arr[c-'1']
	}

	fmt.Println(score)

}
