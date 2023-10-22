package main

import "fmt"

const yes = "YES"
const no = "NO"

func main() {
	var a, b, c, d, j, k int
	fmt.Scanln(&a, &b, &c, &d)

	// if c >= a && d >= b {
	// 	fmt.Println(no)
	// } else {
	// 	k = a / c * b
	// 	j = a/c/2*b + d

	// 	if j < k {
	// 		fmt.Println(yes)
	// 	} else {
	// 		fmt.Println(no)

	// 	}

	// }
	
	// how much time you need to build everything with one hoven
	k = a / c * b
	
	if k > d {
		fmt.Println(no)
	}else if 


	

}
