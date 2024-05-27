package main

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in []int;
		target int;
		want []int
    }{
        {[]int{2,7,11,15}, 9, []int{0,1}},
        {[]int{3,2,4}, 6, []int{1,2}},
        {[]int{3,3}, 6, []int{0,1}},
    }
    for _, c := range cases {
        got := twoSum(c.in, c.target)
		println("got", got)
		if got == nil {
			t.Errorf("Result is nil")
		}
		for i, v := range got {
			if v != c.want[i] {
				t.Errorf("Input %v, target %v, expected indicies %v, but result is %v", c.in, c.target, c.want, got)
			}
		}
    }
}