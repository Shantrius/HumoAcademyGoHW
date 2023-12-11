package main

import "fmt"

func main() {
	var curent_money, transaction int
	var user_name string
	fmt.Scan(&user_name, &curent_money)
	fmt.Println(curent_money)
	fmt.Scan(&transaction)

	if curent_money >= transaction {
		curent_money -= transaction
		fmt.Println(curent_money)
	} else {
		fmt.Println("Недостаточно средств")
	}

}
