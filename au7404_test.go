package main

import (
	"reflect"
	"testing"
)

type testpair7404One struct {
	n, k, fib int
	slice     []int
}

type testpair7404Two struct {
	array [15]float64
	n     int
	ans   [3]float64
}

type testpair7404Three struct {
	n   int
	ans bool
}

type testpair7404Four struct {
	a, b, k int
	res     []int
}

var tests7404One = []testpair7404One{
	{2, 2, 1, []int{1}},
	{8, 8, 21, []int{1, 2, 3, 5, 8, 13, 21}},
}

var tests7404Two = []testpair7404Two{
	{[15]float64{1, 2, 3, 4, 5, 6.087, 7, 8, 9, 10, 11, 12, 13, 14.12, 16}, 15, [3]float64{105.20700000000001, 14.12, 16}},
	{[15]float64{1, 2, 3, 5, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 4, [3]float64{6, 3, 5}},
}

var tests7404Three = []testpair7404Three{
	{1234, false},
	{5321, true},
}

var tests7404Four = []testpair7404Four{
	{1, 10, 4, []int{6, 8, 10}},
	{1, 10, 3, []int{4, 9}},
}

func TestAu7404one(t *testing.T) {
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
	for _, pair := range tests7404Two {
		ans1, ans2, ans3 := Au7404two(pair.array, pair.n)
		result := [3]float64{ans1, ans2, ans3}
		if result != pair.ans {
			t.Error(
				"For", pair.n, pair.array,
				"expected", pair.ans,
				"got", result,
			)
		}
	}
}

func TestAu7404three(t *testing.T) {
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
