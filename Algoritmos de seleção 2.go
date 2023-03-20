package main

import "fmt"

func main() {
	x := 3
	y := 4
	z := 5

	if x < y && x < z {
		fmt.Println("x é menor que y e z")
		if y < x && y < z {
			fmt.Println("y é menor que x e z")
			if z < x && z < y {
				fmt.Println("z é menor que x e y")
			}

		}
	}

}
