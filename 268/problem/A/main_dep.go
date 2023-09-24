package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, s, k, l int
	fmt.Scanln(&n)
	var hs = make([]int, n)
	var as = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&hs[i], &as[i])
	}

	shs := hs[:]
	sas := as[:]

	sort.Ints(shs)
	sort.Ints(sas)

	fmt.Println("___________________________________")
	fmt.Println(shs)
	fmt.Println(sas)
	fmt.Println("___________________________________")

	for i, j := 0, 0; i < n && j < n; {
		if shs[i] == sas[j] {
			k = shs[i]
			l = sas[j]
			s++
			i++
			j++
		} else if shs[i] > sas[j] {
			if k == sas[j] {
				s++
			}
			j++
		} else {
			if l == shs[i] {
				s++
			}
			i++
		}
	}

	fmt.Println(s)

}
