package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		a++
		fmt.Println(a)
	} else {
		var b int = (a - 2)
		fmt.Println(b)
	}
}
