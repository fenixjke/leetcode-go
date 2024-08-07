package task

import "testing"

func TestTwoSums(t *testing.T) {
	cases := []struct {
		in     []int
		target int
		want   []int
	}{
		{InputArrays[0], Targets[0], []int{0, 1}},
		{InputArrays[1], Targets[1], []int{1, 2}},
		{InputArrays[2], Targets[2], []int{0, 1}},
	}
	for _, c := range cases {
		got := twoSum(c.in, c.target)
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
