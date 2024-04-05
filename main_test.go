package main

import "testing"

func TestPowerH(t *testing.T) {
	type test struct {
		Base           int
		Power          int
		ExpectedResult int
	}
	testCases := []test{
		{Base: 0, Power: 1, ExpectedResult: 0},
		{Base: 0, Power: 2, ExpectedResult: 0},
		{Base: 1, Power: 2, ExpectedResult: 1},
		{Base: 2, Power: 3, ExpectedResult: 8},
		{Base: 2, Power: 4, ExpectedResult: 16},
		{Base: 2, Power: 5, ExpectedResult: 32},
		{Base: 2, Power: 6, ExpectedResult: 64},
	}

	for _, cases := range testCases {
		res := PowerH(cases.Base, cases.Power)
		if res != cases.ExpectedResult {
			t.Errorf("expected: %d , result: %d", cases.ExpectedResult, res)
		}
	}
}
