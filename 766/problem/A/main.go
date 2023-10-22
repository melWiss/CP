package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if len(a) > len(b) {
		fmt.Println(len(a))
	} else if len(b) > len(a) {
		fmt.Println(len(b))
	} else if a != b {
		fmt.Println(len(a))
	} else {
		fmt.Println(-1)
	}
}
