package task

import (
	"testing"
)

type TestCase struct {
	i1   []int
	i2   []int
	want float64
}

func Test(t *testing.T) {
	var cases []TestCase
	for i := 0; i < len(input1); i++ {
		cases = append(cases, TestCase{
			i1:   input1[i],
			i2:   input2[i],
			want: expects[i],
		})
	}
	for _, c := range cases {
		got := findMedianSortedArrays(c.i1, c.i2)
		if c.want != got {
			t.Errorf("i1: %v, i2: %v, want: %v, got: %v", c.i1, c.i2, c.want, got)
			t.FailNow()
		}
	}
}
