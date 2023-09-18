package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var s string
	fmt.Scanln(&s)
	var ss = []rune(s)
	var a = 0
	var d = 0
	for i := 0; i < n; i++ {
		if ss[i] == 'A' {
			a++
		} else {
			d++
		}
	}

	if a > d {
		fmt.Println("Anton")
	} else if d > a {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}
}
