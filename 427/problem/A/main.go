package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var entity int
	var nbrOfCrime int
	var nbrOfPolice int

	for i := 0; i < n; i++ {
		fmt.Scan(&entity)
		if entity != -1 {
			nbrOfPolice += entity
		} else {
			if nbrOfPolice > 0 {
				nbrOfPolice--
			} else {
				nbrOfCrime++
			}
		}
	}

	fmt.Println(nbrOfCrime)
}
