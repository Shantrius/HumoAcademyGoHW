package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		d++
	} else {
		e++
	}
	if b > 0 {
		d++
	} else {
		e++
	}
	if c > 0 {
		d++
	} else {
		e++
	}
	fmt.Println(d, e)
}
