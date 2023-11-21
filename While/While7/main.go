package main

import "fmt"

func main() {
	var n int
	var k int = 1
	fmt.Scan(&n)
	for n >= (k * k) {
		k++
	}
	fmt.Println(k)
}
