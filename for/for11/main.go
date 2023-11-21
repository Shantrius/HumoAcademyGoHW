package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		b = i * i
		c += b
	}
	fmt.Println(c)
}
