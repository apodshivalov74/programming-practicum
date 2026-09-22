package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	a := n / 100
	b := (n / 10) % 10
	c := n % 10

	if a != b && a != c && b != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
