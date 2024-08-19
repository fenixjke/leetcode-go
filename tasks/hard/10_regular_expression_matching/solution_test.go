package task

import (
	"testing"
)

type TestCase struct {
	s    string
	p    string
	want bool
}

func Test(t *testing.T) {
	var cases []TestCase
	for i := 0; i < len(input1); i++ {
		cases = append(cases, TestCase{
			s:    input1[i],
			p:    input2[i],
			want: expects[i],
		})
	}
	for _, c := range cases {
		got := isMatch(c.s, c.p)
		if c.want != got {
			t.Errorf("input1: %v, input2: %v, want: %v, got: %v", c.s, c.p, c.want, got)
			t.FailNow()
		}
	}
}
