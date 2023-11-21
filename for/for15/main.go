package main

import "fmt"

func main() {
	var a, b, d int
	fmt.Scan(&a, &d)
	for i := 0; i < d; i++ {
		b *= a
	}
	fmt.Println(b)
}
