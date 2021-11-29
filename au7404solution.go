package main

//	Задача 1
//	Последовательность Фибоначчи образуется так:
//	первый и второй члены последовательности равны 1,
//	каждый следующий равен сумме двух предыдущих (1, 1, 2, 3, 5, 8, 13, ...).
//	Дано натуральное число n (n >= 3).
//	а) Найти k-й член последовательности Фибоначчи.
//	б) Получить первые n членов последовательности Фибоначчи
func Au7404one(n int, k int) (int, []int) {
	var slice = []int{1}

	previousFib := 0
	currentFib := 1

	for i := 1; i < k; i++ {
		newFib := previousFib + currentFib
		if i <= n {
			slice = append(slice, newFib)
		}
		previousFib = currentFib
		currentFib = newFib
	}

	return currentFib, slice
}

// Au7404two Задача 2
// Дана последовательность вещественных чисел a1, a2, ... a15,
// упорядоченная по возрастанию, и число n,
// не равное ни одному из чисел последовательности и такое,
// что a1 < n <a15.
// а) Определить сумму чисел последовательности, меньших n.
// б) Найти два элемента последовательности (их порядковые номера и значение) в интервале,
// между которыми находится значение n
type Answer7404two struct {
	sum                 float64
	firstKey, secondKey int
	firstVal, secondVal float64
}

func Au7404two(array [15]float64, n int) Answer7404two {
	nFloat := float64(n)
	var sum float64

	for k, i := range array {
		if nFloat < i {
			return Answer7404two{sum, k, k + 1, array[k-1], i}
		}
		sum += i
	}
	return Answer7404two{}
}

//	Задача 3
//	Дано натуральное число.
//	Установить, является ли последовательность его цифр при просмотре их справа
//	налево упорядоченной по возрастанию. Например, для числа 5321 ответ положительный, для чисел 7820 и
//	9663 — отрицательный и т. п.
func Au7404three(n int) bool {
	oldDig := 0
	curDig := 0

	for n > 0 {
		oldDig = curDig
		curDig = n % 10
		if curDig < oldDig {
			return false
		}
		n /= 10
	}

	return true
}

// Au7404four Задача 4
//Найти все целые числа из промежутка от a до b, у которых количество делителей равно k.
func Au7404four(a int, b int, k int) []int {
	var result []int
	for i := a; i <= b; i++ {
		del := 1
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				del++
			}
		}
		if del == k {
			result = append(result, i)
		}
	}
	return result
}
