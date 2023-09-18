package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var p, v, t int
	var sum = 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&p, &v, &t)
		if p+v+t >= 2 {
			sum++
		}
	}
	fmt.Println(sum)
}
