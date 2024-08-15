package task

import (
	"testing"
)

type TestCase struct {
	x    string
	want int
}

func Test(t *testing.T) {
	var cases []TestCase
	for i := 0; i < len(input1); i++ {
		cases = append(cases, TestCase{
			x:    input1[i],
			want: expects[i],
		})
	}
	for _, c := range cases {
		got := myAtoi(c.x)
		if c.want != got {
			t.Errorf("input: %v, want: %v, got: %v", c.x, c.want, got)
			t.FailNow()
		}
	}
}
