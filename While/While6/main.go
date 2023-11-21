package main

import "fmt"

func main() {
	var n, divide int
	fmt.Scan(&n)
	divide = n
	if n%2 == 0 {
		for divide > 2 {
			divide -= 2
			n *= divide

		}
	} else {
		for divide > 1 {
			divide -= 2
			n *= divide

		}
	}
	fmt.Println(n)
}
