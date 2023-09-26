package main

import "fmt"

func main() {
	var score int
	var k, i int
	var r int
	fmt.Scanln(&k, &r)

	i = k

	for i%10 != r && i%10 != 0 {
		i = i%10 + k
		score++
	}
	score++
	fmt.Println(score)

}
