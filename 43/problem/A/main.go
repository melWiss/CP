package main

import "fmt"

func main() {
	var n int
	var max int
	var maxTeam string
	var team string
	var game = make(map[string]int)
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&team)
		game[team]++
	}

	for t, s := range game {
		if s > max {
			max = s
			maxTeam = t
		}
	}

	fmt.Println(maxTeam)
}
