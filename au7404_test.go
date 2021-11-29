package main

import (
	"reflect"
	"testing"
)

func TestAu7404one(t *testing.T) {
	var tests7404One = []struct {
		n, k, fib int
		slice     []int
	}{
		{2, 2, 1, []int{1, 1}},
		{8, 8, 21, []int{1, 1, 2, 3, 5, 8, 13, 21}},
	}
	for _, pair := range tests7404One {
		fib, slice := Au7404one(pair.n, pair.k)
		if fib != pair.fib || !reflect.DeepEqual(slice, pair.slice) {
			t.Error(
				"For", pair.n, pair.k,
				"expected", pair.fib, pair.slice,
				"got", fib, slice,
			)
		}
	}
}

func TestAu7404two(t *testing.T) {
	var tests7404Two = []struct {
		array [15]float64
		n     int
		ans   Answer7404two
	}{
		{
			[15]float64{1, 2, 3, 4, 5, 6.087, 7, 8, 9, 10, 11, 12, 13, 14.12, 16},
			15,
			Answer7404two{105.20700000000001, 14, 15, 14.12, 16},
		},
		{
			[15]float64{1, 2, 3, 5, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			4,
			Answer7404two{6, 3, 4, 3, 5},
		},
	}
	for _, pair := range tests7404Two {
		ans := Au7404two(pair.array, pair.n)
		if ans != pair.ans {
			t.Error(
				"For", pair.n, pair.array,
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}

func TestAu7404three(t *testing.T) {
	var tests7404Three = []struct {
		n   int
		ans bool
	}{
		{1234, false},
		{5321, true},
	}
	for _, pair := range tests7404Three {
		ans := Au7404three(pair.n)
		if ans != pair.ans {
			t.Error(
				"For", pair.n,
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}

func TestAu7404four(t *testing.T) {
	var tests7404Four = []struct {
		a, b, k int
		res     []int
	}{
		{1, 10, 4, []int{6, 8, 10}},
		{1, 10, 3, []int{4, 9}},
	}
	for _, pair := range tests7404Four {
		ans := Au7404four(pair.a, pair.b, pair.k)
		if !reflect.DeepEqual(ans, pair.res) {
			t.Error(
				"For", pair.a, pair.b, pair.k,
				"expected", pair.res,
				"got", ans,
			)
		}
	}
}
