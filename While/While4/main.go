package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for {
		n /= 3
		if n < 3 {
			break
		}
	}
	if n == 1 {
		fmt.Print("TRUE")
	} else {
		fmt.Print("FALSE")
	}
}
