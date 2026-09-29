package main

import "fmt"

func main() {
	Task1()
	Task2()
}

//TODO: Задача 1.1
/* Объявить переменные всех основных типов 3 способами:
- var с типом
- var без типа
- :=
Вывести значения через %T
*/

func Task1() {
	var a int = 10
	var b string = "Ivan"
	var c float64 = 3.14
	var d bool = true
	var e byte = 253
	var f rune = 'A'

	var a1 = 20
	var b1 = "Petya"
	var c1 = 3.141
	var d1 = false
	var e1 = 254
	var f1 = 'B'

	a2 := 30
	b2 := "Kolya"
	c2 := 3.1415
	d2 := true
	e2 := 255
	f2 := 'C'

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", e, e)
	fmt.Printf("%v %T\n", f, f)

	fmt.Printf("%v %T\n", a1, a1)
	fmt.Printf("%v %T\n", b1, b1)
	fmt.Printf("%v %T\n", c1, c1)
	fmt.Printf("%v %T\n", d1, d1)
	fmt.Printf("%v %T\n", e1, e1)
	fmt.Printf("%v %T\n", f1, f1)

	fmt.Printf("%v %T\n", a2, a2)
	fmt.Printf("%v %T\n", b2, b2)
	fmt.Printf("%v %T\n", c2, c2)
	fmt.Printf("%v %T\n", d2, d2)
	fmt.Printf("%v %T\n", e2, e2)
	fmt.Printf("%v %T\n", f2, f2)
}


// TODO: Задача 1.2
/* Поменять местами значения двух переменных
- Через третью переменную
- Без третьей переменной
*/

func Task2() {
	var x int = 10
	var y int = 20
	fmt.Printf("x: %v, y: %v\n", x, y)

	var z int = x
	x = y
	y = z
	fmt.Printf("x: %v, y: %v\n", x, y)	

	var x1 int = 10;
	var y1 int = 20
	x1, y1 = y1, x1
	fmt.Printf("x1: %v, y1: %v\n", x1, y1)
}
