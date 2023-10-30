package main

import "fmt"

func main() {
	var a, b, totalHappyNumbers int
	fmt.Scanln(&a, &b)
	var numbers = make(map[int]int)

	for n := a; n <= b; n++ {
		for i := calc(n); ; i = calc(i) {
			numbers[i] += 1
			if numbers[i] >= 2 {
				if i == 1 {
					fmt.Println(n)
					totalHappyNumbers++
				}
				break
			}
		}

		numbers = make(map[int]int)
	}

	fmt.Println(totalHappyNumbers)

}

func calc(n int) int {
	k := n
	sum := 0
	for k = n % 10; k > 0 || n > 0; k = n % 10 {
		sum += k * k
		n /= 10

	}

	return sum
}
