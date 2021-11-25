package main

import (
	"fmt"
)

//задача 2
var h int
var m int
var s int

//задача 3
var integer int
var testInteger int

//задача 4
var yBirth int
var mBirth int

func main() {
	fmt.Println("Задача 1. Результат для числа 237: ", TaskOne())

	fmt.Println("Задача 2. Введите часы, минуты и секунды: ")
	fmt.Scan(&h, &m, &s)
	fmt.Println("Задача 2. Угол равен ", TaskTwo(h, m, s))

	fmt.Println("Задача 3. Введите 4-х значное число и число для проверки кратности: ")
	fmt.Scan(&integer, &testInteger)
	a1, a2, a3, a4 := TaskThree(integer, testInteger)
	fmt.Printf(
		`Задача 3.
	1) равна ли сумма двух первых его цифр сумме двух его последних цифр - %t
	2) кратна ли трем сумма его цифр - %t
	3) кратно ли четырем произведение его цифр - %t
	4) кратно ли произведение его цифр числу а - %t`,
		a1, a2, a3, a4,
	)

	fmt.Println("\r\n Задача 4. Введите год (4 цифры) и месяц рождения (число, без 0): ")
	fmt.Scan(&yBirth, &mBirth)
	year, month := TaskFour(yBirth, mBirth)
	fmt.Printf("\r\n Задача 4. Сейчас человеку полных лет %d, полных месяцев %d", year, month)
}
