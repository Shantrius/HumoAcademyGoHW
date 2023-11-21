package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(a)
	for i := 0; i < a; i++ {
		b += 1 / (i + 1)
	}
	fmt.Println(b)
}
