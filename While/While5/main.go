package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)
	if n%2 == 0 {
		for n >= 2 {
			n /= 2
			count++
		}
		fmt.Println(count)
	}
}
