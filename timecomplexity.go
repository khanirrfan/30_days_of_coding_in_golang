package main

import (
	"fmt"
	"math"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t int
	fmt.Scan(&t)

	n := make([]float64, t)

	for k := range n {
		fmt.Scan(&n[k])
	}
	for _, value := range n {
		fmt.Println(checkPrimeNumber(value))
	}
}

func checkPrimeNumber(num float64) string {
	var i int
	if num == 1 || (num != 2 && int(num)%2 == 0) {
		return "Not prime"
	}
	for i = 3; i <= int(math.Sqrt(num)); i += 2 {
		if int(num)%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}
