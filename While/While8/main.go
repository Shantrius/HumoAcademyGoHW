package main

import "fmt"

func main() {
	var n, k int
	k = 1
	fmt.Scan(&n)
	for n >= (k * k) {
		k++
	}
	fmt.Println(k - 1)
}
