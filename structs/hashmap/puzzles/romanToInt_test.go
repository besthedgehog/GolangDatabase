package main

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Errorf("III should be equal to 3")
	}
	if romanToInt("LVIII") != 58 {
		t.Errorf("LVIII should be equal to 58")
	}
	if romanToInt("MCMXCIV") != 1994 {
		t.Errorf("MCMXCIV should be equal to 1994, but val = %v", romanToInt("MCMXCIV"))
	}
}
