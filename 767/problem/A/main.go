package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, i, j int
	fmt.Scanln(&n)
	arr := make([]int, n)
	potato := make(map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		potato[i+1] = false
	}
	arrSorted := make([]int, n)
	copy(arrSorted, arr)
	arrSortedSlice := arrSorted[:]
	sort.Sort(sort.Reverse(sort.IntSlice(arrSortedSlice)))

	for i = 0; i < n; {
		for j = i; j < n; j++ {
			if arrSorted[i] != arr[j] {
				if !potato[arr[j]] {
					fmt.Println("first", j)
					potato[arr[j]] = true
				}
			} else {
				j++
				potato[arr[j]] = true
				break
			}
		}
		for ; i < j; i++ {
			if potato[arrSorted[i]] {
				fmt.Printf("%d ", arrSorted[i])
			} else {
				break
			}
		}
		fmt.Println("second ", i)
	}

}
