package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	for n >= 10 {
		n = n / 10
	}

	fmt.Println(n)
}
