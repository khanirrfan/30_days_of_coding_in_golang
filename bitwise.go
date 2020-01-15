package main

import (
	"fmt"
)

func main() {
	var t, n, k int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		fmt.Scan(&k)
		if ((k - 1) | k) <= n {
			fmt.Println(k - 1)
		} else {
			fmt.Println(k - 2)
		}
	}

}
