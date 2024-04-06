package main

import "testing"

func TestPowerH(t *testing.T) {
	type test struct {
		Base           int
		Power          int
		ExpectedResult int
	}

	var testCases = []test{
		{Base: 1, Power: 2, ExpectedResult: 1},
		{Base: 0, Power: 2, ExpectedResult: 0},
		{Base: 10, Power: 2, ExpectedResult: 100},
		{Base: 1, Power: 200, ExpectedResult: 1},
		{Base: 1, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 400, Power: 2, ExpectedResult: 160000},
		{Base: 0, Power: 0, ExpectedResult: 1},
	}

	for _, c := range testCases {
		res := PowerH(c.Base, c.Power)

		if res != c.ExpectedResult {
			t.Errorf("expecte: %d , result: %d", c.ExpectedResult, res)
		}
	}

}

var result int

func BenchmarkPowerH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PowerH(2, 1)
		result = res
	}
}
