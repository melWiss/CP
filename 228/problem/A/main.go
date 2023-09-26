package main

import "fmt"

func main() {
	var s1, s2, s3, s4, result int
	var colors = make(map[int]int)

	fmt.Scanln(&s1, &s2, &s3, &s4)

	colors[s1] = s1
	colors[s2] = s2
	colors[s3] = s3
	colors[s4] = s4

	result = 4 - len(colors)

	fmt.Println(result)

}
