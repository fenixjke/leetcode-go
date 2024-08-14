package task

import (
	"testing"
)

type TestCase struct {
	s       string
	numRows int
	want    string
}

func Test(t *testing.T) {
	var cases []TestCase
	for i := 0; i < len(input1); i++ {
		cases = append(cases, TestCase{
			s:       input1[i],
			numRows: input2[i],
			want:    expects[i],
		})
	}
	for _, c := range cases {
		got := convert(c.s, c.numRows)
		if c.want != got {
			t.Errorf("str: %v, want: %v, got: %v", c.s, c.want, got)
			t.FailNow()
		}
	}
}
