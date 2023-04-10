package main

import "testing"

// Табличный тест
var tests = []struct {
	input16  string
	output10 string
}{
	{"", ""},
	{"174214D1D6ADAC", "6546581651631532"},
	{"174214D1D6ADAB", "6546581651631531"},
	{"1", "1"},
	{"2386F26FC10000", "10000000000000000"},
}

func TestConvertHEX2DEC(t *testing.T) {
	for _, e := range tests {
		res := ConvertHEX2DEC(e.input16)
		if res != e.output10 {
			t.Errorf("Получено: %v, Ожидалось: %v, Ввод: %v", res, e.output10, e.input16)
		}
	}
}
