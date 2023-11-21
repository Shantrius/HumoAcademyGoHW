package main

import "fmt"

func main() {
	var n, k, count int
	k = 3
	count = 1
	fmt.Scan(&n)
	for n > k {
		k *= 3
		count++
	}
	fmt.Println(count)
}
