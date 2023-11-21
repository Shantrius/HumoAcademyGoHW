package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b)
	if a > b {
		c = a
		d = b
	} else {
		c = b
		d = a
	}
	fmt.Println(c, d)
}
