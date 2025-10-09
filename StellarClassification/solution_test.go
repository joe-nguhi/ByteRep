package main

import "testing"

func TestClassification(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{5778, "G"},
		{2400, "M"},
		{9999, "A"},
		{3700, "K"},
		{3699, "M"},
		{210000, "O"},
		{6000, "F"},
		{11432, "B"},
	}

	for _, test := range tests {
		if r := classification(test.input); r != test.want {
			t.Errorf("classification(%d)=%v, wants %s", test.input, r, test.want)
		}
	}

}
