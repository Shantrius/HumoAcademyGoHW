package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println(((a % 2) == 0) || ((b % 2) == 0))
}
