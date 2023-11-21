package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		d++
	}
	if b > 0 {
		d++
	}
	if c > 0 {
		d++
	}
}
