package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	var n int
	var name, email string
	var names []string
	fmt.Scan(&n)
	r, _ := regexp.Compile(`(@gmail\.com)$`)
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scan(&email)
		if r.MatchString(email) {
			names = append(names, name)
		}
	}
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	for _, value := range names {
		fmt.Println(value)
	}
}
