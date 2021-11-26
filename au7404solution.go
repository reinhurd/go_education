package main

//	Задача 1
//	Последовательность Фибоначчи образуется так:
//	первый и второй члены последовательности равны 1,
//	каждый следующий равен сумме двух предыдущих (1, 1, 2, 3, 5, 8, 13, ...).
//	Дано натуральное число n (n >= 3).
//	а) Найти k-й член последовательности Фибоначчи.
//	б) Получить первые n членов последовательности Фибоначчи
func Au7404one(n int, k int) (int, []int) {
	var slice []int

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
func Au7404two(array [15]float64, n int) (float64, float64, float64) {
	nFloat := float64(n)
	key := binarySearch(array, nFloat)
	var sum float64
	for _, i := range array[:key+1] {
		sum += i
	}

	return sum, array[key], array[key+1]
}

func binarySearch(array [15]float64, nFloat float64) int {
	const countAr = 15
	maxKey := countAr - 1

	for lowKey := 0; lowKey <= maxKey; {
		mid := (lowKey + maxKey) / 2
		guess := array[mid]

		if nFloat > array[mid] && nFloat < array[mid+1] {
			return mid
		}
		if guess < nFloat {
			lowKey = mid + 1
		} else {
			maxKey = mid - 1
		}
	}
	return 0
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
		del := 0
		for j := 1; j <= i; j++ {
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
