package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		c = c * i

	}
	fmt.Println(c)
}
