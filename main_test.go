package main

import (
	"testing"
)

type testpairTwo struct {
	h, m, s, result int
}

type testpairThree struct {
	i, a    int
	results [4]bool
}

type testpairFour struct {
	y, m    int
	results [2]int
}

var testsTwo = []testpairTwo{
	{12, 0, 0, 0},
	{24, 59, 59, 29},
	{6, 0, 0, 180},
	{15, 8, 45, 94},
}

var testsThree = []testpairThree{
	{1111, 4, [4]bool{true, false, false, false}},
	{4444, 4, [4]bool{true, false, true, true}},
}

//тестирование функций, которые берут дату из системного времени по жестко заданному датапровайдеру
// - вообще плохая практика
//но для текущего задания представим, что тесты будут запускаться только в ноябре)
var testsFour = []testpairFour{
	{1987, 11, [2]int{34, 0}},
	{2000, 4, [2]int{21, 7}},
	{2021, 11, [2]int{0, 0}},
	{2020, 12, [2]int{0, 11}},
}

func TestTaskOne(t *testing.T) {
	const mockNumber = 237
	const correctAnswer = 372
	v := TaskOne()
	if v != correctAnswer {
		t.Error(
			"For", mockNumber,
			"expected", correctAnswer,
			"got", v,
		)
	}
}

func TestTaskTwo(t *testing.T) {
	for _, pair := range testsTwo {
		v := TaskTwo(pair.h, pair.m, pair.s)
		if v != pair.result {
			t.Error(
				"For", pair.h, pair.m, pair.s,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestTaskThree(t *testing.T) {
	for _, pair := range testsThree {
		a1, a2, a3, a4 := TaskThree(pair.i, pair.a)
		results := [4]bool{a1, a2, a3, a4}
		if results != pair.results {
			t.Error(
				"For", pair.i, pair.a,
				"expected", pair.results,
				"got", results,
			)
		}
	}
}

func TestTaskFour(t *testing.T) {
	for _, pair := range testsFour {
		yR, mR := TaskFour(pair.y, pair.m)
		results := [2]int{yR, mR}
		if results != pair.results {
			t.Error(
				"For", pair.y, pair.m,
				"expected", pair.results,
				"got", results,
			)
		}
	}
}
