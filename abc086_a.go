package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	product := a * b
	remainder := product % 2
	if remainder == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
