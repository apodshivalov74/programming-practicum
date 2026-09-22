package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	d1 := n / 100000
	d2 := (n / 10000) % 10
	d3 := (n / 1000) % 10
	d4 := (n / 100) % 10
	d5 := (n / 10) % 10
	d6 := n % 10

	sumFirst := d1 + d2 + d3
	sumLast := d4 + d5 + d6

	if sumFirst == sumLast {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
