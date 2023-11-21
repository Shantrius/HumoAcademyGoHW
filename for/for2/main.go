package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(a, b)
	for i := 0; i <= b; i++ {
		var c int
		c = a
		fmt.Print(c)
		c++
	}
}
