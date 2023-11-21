package main

import "fmt"

func main() {
	var a, b float64
	b = 1.2
	fmt.Scan(&a)
	for i := 0; i < 5; i++ {
		fmt.Println(a * b)
		b += 0.2
	}
}
