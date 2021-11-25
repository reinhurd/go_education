package main

import (
	"strconv"
	"time"
)

// TaskOne Задача 1
//Из трехзначного числа x вычли его последнюю цифру.
//Когда результат разделили на 10, а к частному слева приписали последнюю цифру числа x,
//то получилось число 237. Написать функцию для нахождения числа x.
func TaskOne() int {
	res := 237
	resString := strconv.Itoa(res)
	firstDigit, _ := strconv.Atoi(resString[:1])
	origIntWithoutLastDigit, _ := strconv.Atoi(resString[1:] + "0")
	return origIntWithoutLastDigit + firstDigit
}

// TaskTwo Задача 2
//Даны целые числа h, m, s (0 < h ≤ 23, 0 ≤ m ≤ 59, 0 ≤ s ≤ 59),
//указывающие момент времени: "h часов, m минут, s секунд".
//Написать функцию для определения угла (в градусах) между положением часовой стрелки
//в начале суток и в указанный момент времени
func TaskTwo(h int, m int, s int) int {
	const secInHour = 3600
	const minInHour = 60
	const fullCircleDeg = 360
	const fullCircleHour = 12

	hHalfDay := h % fullCircleHour
	sFull := hHalfDay * secInHour + m * minInHour + s
	res := fullCircleDeg * sFull / fullCircleHour / secInHour
	return res
}

// TaskThree Задача 3
//Дано четырехзначное число.
//Написать функцию для определения:
//1) равна ли сумма двух первых его цифр сумме двух его последних цифр;
//2) кратна ли трем сумма его цифр;
//3) кратно ли четырем произведение его цифр;
//4) кратно ли произведение его цифр числу а.
func TaskThree(i int, a int) (ans1 bool, ans2 bool, ans3 bool, ans4 bool) {
	s := strconv.Itoa(i)
	i1, _ := strconv.Atoi(s[0:1])
	i2, _ := strconv.Atoi(s[1:2])
	i3, _ := strconv.Atoi(s[2:3])
	i4, _ := strconv.Atoi(s[3:4])

	ans1 = (i1 + i2) == (i3 + i4)
	ans2 = (i1 + i2 + i3 + i4) % 3 == 0
	ans3 = (i1 * i2 * i3 * i4) % 4 == 0
	ans4 = (i1 * i2 * i3 * i4) % a == 0
	return
}

// TaskFour Задача 4
//Известны год и номер месяца рождения человека,
//а также год и номер месяца сегодняшнего дня.
//Определить возраст человека (число полных лет и число полных месяцев).
//При определении числа полных месяцев дни месяца не учитывать,
//а использовать разность между номерами месяцев. Например,
//если месяц рождения — февраль, а текущий (сегодняшний) месяц — май,
//то число полных месяцев равно трем независимо от дня рождения и сегодняшнего дня.
//
//не совсем ясно из задачи, требуются ли кастомные "сегодняшние даты"
//сделал их автоподгрузку из системных
func TaskFour(yMan int, mMan int) (yFull int, mFull int) {
	const monthInYear = 12
	year, month, _ := time.Now().Date()
	monthInt := int(month)

	//Если ДР уже случилось в этом году
	if monthInt >= mMan {
		yFull = year - yMan
		mFull = monthInt - mMan
	} else {
		yFull = year - yMan - 1
		mFull = monthInYear - mMan + monthInt
	}
	return
}
