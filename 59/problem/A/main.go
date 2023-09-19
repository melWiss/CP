package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	chars := []rune(s)
	up := 0
	low := 0

	for _, c := range chars {
		if c <= 'Z' {
			up++
		} else {
			low++
		}
	}

	if up > low {
		fmt.Println(strings.ToUpper(string(chars)))
	} else {
		fmt.Println(strings.ToLower(string(chars)))

	}

}
