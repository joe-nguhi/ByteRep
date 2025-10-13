package main

import "testing"

func TestMoonPhase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2000-01-12",
			want:  "New",
		},
		{
			input: "2000-01-13",
			want:  "Waxing",
		},
		{
			input: "2014-10-15",
			want:  "Full",
		},
		{
			input: "2012-10-21",
			want:  "Waning",
		},
		{
			input: "2022-12-14",
			want:  "New",
		},
	}

	for _, test := range tests {
		if result := moonPhase(test.input); result != test.want {
			t.Errorf("moonPhase(%s) = %s, want %s", test.input, result, test.want)
		}
	}
}
