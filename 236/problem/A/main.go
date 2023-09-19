package main

import "fmt"

func main() {
	var name string
	fmt.Scan(&name)

	chars := []rune(name)

	charset := make(map[rune]rune, len(chars))

	for _, c := range chars {
		charset[c] = c
	}

	r := len(charset)

	if r%2 != 0 {
		fmt.Println("IGNORE HIM!")
	} else {
		fmt.Println("CHAT WITH HER!")
	}
}
