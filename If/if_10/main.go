package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	if !(a == b) {
		c = (a + b)
		a = c
		b = c
	} else {
		a = 0
		b = 0
	}
	fmt.Println(a, b)

}
