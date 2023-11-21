package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		b = 2*i - 1
		c += b
	}
	fmt.Println(c)
}
