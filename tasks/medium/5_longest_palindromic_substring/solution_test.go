package task

import (
	"testing"
)

type TestCase struct {
	s    string
	want string
}

func TestAddTwoNumbers(t *testing.T) {
	var cases []TestCase
	for i := 0; i < len(input); i++ {
		cases = append(cases, TestCase{
			s:    input[i],
			want: expects[i],
		})
	}
	for _, c := range cases {
		got := lengthOfLongestSubstring(c.s)
		if c.want != got {
			t.Errorf("str: %v, want: %v, got: %v", c.s, c.want, got)
			t.FailNow()
		}
	}
}
