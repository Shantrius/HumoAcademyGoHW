package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		a++
		fmt.Println(a)
	}

}
