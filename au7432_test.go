package main

import (
	"reflect"
	"testing"
)

func TestAu7432one(t *testing.T) {
	var testData = []struct {
		s   string
		ans int
	}{
		{"GG", 2},
		{"ППрпп", 4},
		{"ППрппп", 5},
		{"ППрппп ggHHiutyGGGgg", 14},
		{"1ggg_o__", 5},
	}
	for _, pair := range testData {
		ans := Au7432one(pair.s)
		if ans != pair.ans {
			t.Error(
				"For", pair.s,
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}

func TestAu7432two(t *testing.T) {
	var testData = []struct {
		words [3]string
		ans   []string
	}{
		{
			[3]string{"rТТапрр", "GGgiTT", "ААрваЛл"},
			[]string{"rап", "gi", "рваЛл"},
		},
	}
	for _, pair := range testData {
		ans := Au7432two(pair.words)
		if !reflect.DeepEqual(ans, pair.ans) {
			t.Error(
				"For", pair.words,
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}

func TestAu7432three(t *testing.T) {
	var testData = []struct {
		sentences [2]string
		ans       []string
	}{
		{
			[2]string{
				"раз два три три четыре",
				"два три четыре четыре пять",
			},
			[]string{"пять", "раз"},
		},
		{
			[2]string{
				"раз",
				"два три",
			},
			[]string{"два", "раз", "три"},
		},
	}
	for _, pair := range testData {
		ans := Au7432three(pair.sentences)
		if !reflect.DeepEqual(ans, pair.ans) {
			t.Error(
				"For", pair.sentences,
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}

//тестирование i/o в юнитах - это плохо, поэтому тут просто проверка записываемого результата
func TestAu7432four(t *testing.T) {
	var testData = []struct {
		ans []string
	}{
		{
			[]string{
				"1234575869679326547964327856329562395629835832659438568235",
				"ghfjkgjhsdkfgkystytoutyithjfghsdlghfdklghjksfdhgflkdshksfghlk",
				"пароплварыпывдапавыолпрыдвапы4",
			},
		},
	}
	for _, pair := range testData {
		ans := Au7432four()
		if !reflect.DeepEqual(ans, pair.ans) {
			t.Error(
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}

//аналогично с тестированием пятой задачи
func TestAu7432five(t *testing.T) {
	var testData = []struct {
		ans int
	}{
		{2},
	}
	for _, pair := range testData {
		ans := Au7432five()
		if ans != pair.ans {
			t.Error(
				"expected", pair.ans,
				"got", ans,
			)
		}
	}
}
