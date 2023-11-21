package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	for a >= b {
		c++
	}
	fmt.Println(c)
}
