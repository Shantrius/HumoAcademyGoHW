package main

import "fmt"

func main() {
	var a, b, d int
	fmt.Scan(a, b)
	for i := 0; i < b; i++ {
		var c, d int
		c = b - 1

		fmt.Print(c)
		c--
		d++
	}
	fmt.Print(d)
}
