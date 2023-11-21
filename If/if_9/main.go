package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a, &b)
	if a > b {

		a = (a + b)
		b = (a - b)
		a = (a - b)
		fmt.Println(a, b)
	} else {
		fmt.Println(a, b)
	}

}
