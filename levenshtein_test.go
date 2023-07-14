package levenshtein

import (
	"testing"
)

type TestCase struct {
	S1 string
	S2 string
	R  int
}

func TestMin(t *testing.T) {
	if r := min(2, 1, 3); r != 1 {
		t.Error("wrong")
	}
	if r := min(9, 8, 7); r != 7 {
		t.Error("wrong")
	}
	if r := min(222, 30, 100); r != 30 {
		t.Error("wrong")
	}
}

func TestLevenshtein(t *testing.T) {

	testCases := []TestCase{
		TestCase{
			"res",
			"eys",
			2,
		},
		TestCase{
			"me",
			"my",
			1,
		},
		TestCase{
			"pferd",
			"apfel",
			3,
		},
		TestCase{
			"",
			"",
			0,
		},
		TestCase{
			"foo",
			"foo",
			0,
		},
		TestCase{
			"foo",
			"bar",
			3,
		},
	}

	for _, testCase := range testCases {
		res := Distance(testCase.S1, testCase.S2)
		if res != testCase.R {
			t.Errorf("wrong - %v %v %d should be %d", testCase.S1, testCase.S2, res, testCase.R)
		}
	}

}
