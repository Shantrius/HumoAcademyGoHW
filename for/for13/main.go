package main

import "fmt"

func main() {
	var a, b, c float64
	var d int
	fmt.Scan(&a)
	d = int(a)
	for i := 0; i < d; i++ {
		b = (a + float64(i)*0.1) * (float64(i) + 1)
		c += b
	}
	fmt.Println(c)
}
