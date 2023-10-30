package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	var a float64
	// fmt.Scan(&t)
	t = 1
	for i := 0; i < t; i++ {
		fmt.Scan(&a)
		if a < 0 {
			a += 360
		}

		goal := 0.0

		for goal <= 90 {
			n1 := (90 - goal) / a
			n2 := (270 + goal) / a

			if n1 != float64(int64(n1)) {
				if n2 != float64(int64(n2)) {
					goal++
					continue
				}
				fmt.Println(n2)
			} else {
				if n1 > 0 {
					fmt.Println(n1)
				}
			}
			break
		}

		if goal >= 90 {
			angle := 450.0
			n := angle / a
			for n != float64(int64(n)) {
				angle += 180
				n = angle / a
			}
			fmt.Println(int64(n))
			fmt.Printf("cos(n*a) = cos(%f, %f)= cos(%f) = %f\n", n, a, n*a, math.Cos(n*a))

		}
	}

}
