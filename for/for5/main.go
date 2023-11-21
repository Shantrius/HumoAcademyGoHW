package main

import "fmt"

func main() {
	var a, b float64
	b = 0.1
	fmt.Scan(&a)
	for i := 0; i < 10; i++ {
		fmt.Println(a * b)
		b += 0.1
	}
}
