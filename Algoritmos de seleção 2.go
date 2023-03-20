package main

import "fmt"

func main() {
	var num int
	fmt.Println("digite um número")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("o número é par")

	} else {
		fmt.Println("o número é ímpar")
	}

}