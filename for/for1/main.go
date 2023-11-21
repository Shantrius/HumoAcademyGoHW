package main

import "fmt"

func main() {
	var N int
	var K int
	fmt.Scan(&N, &K)
	for i := 0; i < N; i++ {
		fmt.Print(K)
	}
}
