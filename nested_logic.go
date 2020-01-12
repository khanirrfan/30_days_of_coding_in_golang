package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var da, ma, ya int64
	var de, me, ye int64
	var fine, late int64
	fmt.Scan(&da)
	fmt.Scan(&ma)
	fmt.Scan(&ya)
	fmt.Scan(&de)
	fmt.Scan(&me)
	fmt.Scan(&ye)
	if ya <= ye {
		if ma <= me {
			if da <= de {
				fine = 0
				fmt.Println(fine)
			} else if da >= de {
				late = da - de
				fine = 15 * late
				fmt.Println(fine)
			}
		} else if ma > me {
			if ya < ye {
				late = ma - me
				fine = 0
				fmt.Println(fine)
			} else if ya == ye {
				late = ma - me
				fine = 500 * late
				fmt.Println(fine)
			}
		}
	} else {
		fine = 10000
		fmt.Println(fine)
	}
}
