package main

import (
	"fmt"
	"strings"
)

func main() {
	var first, second string
	fmt.Scanln(&first)
	fmt.Scanln(&second)

	first = strings.ToLower(first)
	second = strings.ToLower(second)

	// convert to runes
	var firstRunes, secondRunes []rune
	firstRunes = []rune(first)
	secondRunes = []rune(second)

	result := 0

	for i := 0; i < len(firstRunes); i++ {
		result += int(firstRunes[i]) - int(secondRunes[i])
		if result != 0 {
			break
		}
	}

	if result > 0 {
		fmt.Println(1)
	} else if result < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(0)
	}

}
