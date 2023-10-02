package main

import "fmt"

func main() {
	var y, w int
	fmt.Scanln(&y, &w)

	var max int
	var score int

	if y >= w {
		max = y
	} else {
		max = w
	}

	score = 7 - max

	if score == 6 {
		fmt.Println("1/1")
	} else {
		var dom int = 6
		if score%3 == 0 && dom%3 == 0 {
			score /= 3
			dom /= 3
		}

		if score%2 == 0 && dom%2 == 0 {
			score /= 2
			dom /= 2
		}

		fmt.Printf("%d/%d\n", score, dom)
	}

}
