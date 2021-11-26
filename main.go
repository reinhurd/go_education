package main

import "fmt"

func main() {
	//au7403()
	au7404()
}

func au7403() {
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

func au7404() {
	var k int
	var n int
	fmt.Println("Задача 1. Введите порядковый номер числа и количество чисел в массиве для вывода:")
	fmt.Scan(&k, &n)
	ans, slice := Au7404one(n, k)
	fmt.Printf("\r\n Задача 1. Ответ. Число равно %d, массив чисел %v", ans, slice)

	fmt.Println("Задача 2")
	//для этой задачи не стал делать ввод, 15 чисел вводить через консоль такое себе
	arr := [15]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16}
	n2 := 15
	i1, i2, i3 := Au7404two(arr, n2)
	fmt.Printf("\r\n Задача 2. Ответ. Сумма последовательности %f, два чиcла рядом %f, %f", i1, i2, i3)

	fmt.Println("Задача 3. Введите число")
	var n3 int
	fmt.Scan(&n3)
	fmt.Print("Задача 3, ответ: ", Au7404three(n3))

	fmt.Println("Задача 4. Введите первое, последнее число и количество делителей")
	var a4, b4, n4 int
	fmt.Scan(&a4, &b4, &n4)
	fmt.Print("Задача 4. Ответ", Au7404four(a4, b4, n4))
}
