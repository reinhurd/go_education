package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Au7432one Задача 1
//Дано предложение. Определить, сколько в нем одинаковых соседних букв.
//Прим. как я понял, считаем кейсы повторов без уникальности, например "ППрпп" = 4
func Au7432one(s string) int {
	runes := []rune(s)
	charSum := 0

	for k, i := range runes {
		//поскольку считаем каждую следующую букву, нужно сделать отдельный чек самой первой в начале строки
		//к сожалению, не как в пхп, нужно явно контролировать ключ в слайсе
		if k == 0 {
			if i == runes[k+1] {
				charSum++
			}
		} else {
			//чек буквы, на которой начинается ряд напр "1ggg_o" = 3
			if i != runes[k-1] && len(runes) > (k+1) && i == runes[k+1] {
				charSum++
				//чек обычной буквы в середине и конце ряда
			} else if i == runes[k-1] {
				charSum++
			}
		}
	}
	return charSum
}

// Au7432two Задача 2
//Даны три слова. Напечатать неповторяющиеся в них буквы.
//формулировка очень неясна, представим, что на вход три слова как строки
//а неповторяющиеся - буквы, которые не идут друг за другом без проверки на уникальность
func Au7432two(words [3]string) []string {
	var r []string

	//алгоритм обратен предыдущей задаче
	for _, s := range words {
		result := ""
		runes := []rune(s)
		for k, i := range runes {
			if k == 0 {
				if i != runes[k+1] {
					result += string(i)
				}
			} else {
				if i != runes[k-1] && len(runes) > (k+1) && i != runes[k+1] {
					result += string(i)
				} else if k == len(runes)-1 && i != runes[k-1] {
					result += string(i)
				}
			}
		}
		r = append(r, result)
	}
	return r
}

// Au7432three Задача 3
//Даны два предложения. Напечатать слова, которые встречаются в двух предложениях только один раз.
//опять же, условия неясны, представим, что нужно найти уникальные слова в обоих предложениях вместе
func Au7432three(incS [2]string) []string {
	resMap := make(map[string]int)
	var res []string

	for _, s := range incS {
		ar := strings.Fields(s)
		for _, i := range ar {
			resMap[i]++
		}
	}

	for k, i := range resMap {
		if i == 1 {
			res = append(res, k)
		}
	}

	//поскольку при использовании мапы нам не гарантирован порядок,
	//тесты (через DeepEqual) плавают и поэтому нужна дополнительная сортировка
	//альтернативно в тестах можно было бы слайсы проверять по наполнению, а не через DeepEqual
	sort.Strings(res)

	return res
}

// Au7432four Задача 4
//Имеется текстовый файл. Переписать в другой файл те его строки, в которых
//имеется более 30-ти символов
func Au7432four() []string {
	var text []string //строки длиннее 30
	oldfileName := "oldfile.txt"
	newfileName := "newfile.txt"

	oldfile, err := os.Open(oldfileName)
	if err != nil {
		log.Fatalf("failed to open")
	}
	defer oldfile.Close()
	newfile, err := os.Create(newfileName)
	if err != nil {
		fmt.Println("failed to create")
	}
	defer newfile.Close()

	scanner := bufio.NewScanner(oldfile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(scanner.Text()) > 30 {
			text = append(text, scanner.Text())
		}
	}

	w := bufio.NewWriter(newfile)

	for _, line := range text {
		w.WriteString(line + "\n")
	}
	w.Flush()

	return text
}

// Au7432five Задача 5
//Имеются два текстовых файла с одинаковым числом строк. Выяснить, совпадают ли их строки. Если нет, то получить номер первой строки, в которой
//эти файлы отличаются друг от друга.
//возьмем те же файлы, что и в четвертой задаче
func Au7432five() int {
	oldfileName := "oldfile.txt"
	newfileName := "newfile.txt"
	lineCounter := 1

	oldfile, err := os.Open(oldfileName)
	if err != nil {
		log.Fatalf("failed to open")
	}
	defer oldfile.Close()
	newfile, err := os.Open(newfileName)
	if err != nil {
		fmt.Println("failed to open")
	}
	defer newfile.Close()

	scanner1 := bufio.NewScanner(oldfile)
	scanner2 := bufio.NewScanner(newfile)
	scanner1.Split(bufio.ScanLines)
	scanner2.Split(bufio.ScanLines)
	for scanner1.Scan() {
		scanner2.Scan()
		if scanner1.Text() != scanner2.Text() {
			return lineCounter
		}
		lineCounter++
	}

	return lineCounter
}
