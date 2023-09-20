package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	var arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	s := 0
	d := 0
	i := 0
	j := n - 1

	switcher := true

	for i <= j {
		if switcher {
			if arr[i] > arr[j] {
				s += arr[i]
				i++
			} else {
				s += arr[j]
				j--
			}
		} else {
			if arr[i] > arr[j] {
				d += arr[i]
				i++
			} else {
				d += arr[j]
				j--
			}
		}
		switcher = !switcher
	}

	fmt.Println(s, d)
}
