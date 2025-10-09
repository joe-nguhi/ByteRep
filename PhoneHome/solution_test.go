package main

import "testing"

func TestSendMessage(t *testing.T) {
	tests := []struct {
		route []int
		want  string
	}{
		{[]int{300_000, 300_000}, "2.5"},
		{[]int{384_400, 384_400}, "3.0627"},
		{[]int{54_600_000, 54_600_400}, "364.5"},
		{[]int{1_000_000, 500_000_000, 1_000_000}, "1674.3333"},
		{[]int{10_000, 21_339, 50_000, 31_243, 10_000}, "2.4086"},
		{[]int{802_101, 725_994, 112_808, 3_625_770, 481_239}, "21.1597"},
	}

	for _, test := range tests {
		if result := sendMessage(test.route); result != test.want {
			t.Errorf("sendMessage(%v)=%v, wants%v", test.route, result, test.want)
		}
	}
}
