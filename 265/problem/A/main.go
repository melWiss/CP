package main

import "fmt"

func main() {
	var i int
	var s, t string
	fmt.Scanln(&s)
	fmt.Scanln(&t)

	var arrs = []rune(s)
	var arrt = []rune(t)

	for j := 0; j < len(t) && i < len(s); j++ {
		if arrs[i] == arrt[j] {
			i++
		}
	}

	i++

	fmt.Println(i)

}
