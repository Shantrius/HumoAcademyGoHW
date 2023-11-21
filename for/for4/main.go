package main

import "fmt"

func main() {
	var a float64
	var b float64
	fmt.Scan(&a)
	for i := 0; i < 10; i++ {
		fmt.Println(a * b)
		b++
	}
}
