package main

import "fmt"

func main() {
	var n, k, result int
	fmt.Scan(&n, &k)
	if n >= k {
		for n >= k {
			result = n - k
		}
	} else {
		for k >= n {
			result = k - n
		}
	}
	fmt.Println(result)
}
