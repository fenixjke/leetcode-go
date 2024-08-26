package task

import (
	"leetcode-go/common"
)

type S struct{}

var option string = "my"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my", "leetcode"},
		TaskName: "Container With Most Water",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = maxArea(input1[i])
	}
}

func maxArea(height []int) int {
	switch option {
	case "my":
		cap := 0
		p1, p2 := 0, len(height)-1
		for p1 < p2 {
			t := min(height[p1], height[p2]) * (p2 - p1)
			if t > cap {
				cap = t
			}
			if height[p1] < height[p2] {
				p1++
			} else {
				p2--
			}
		}
		return cap
	case "leetcode":
		size := len(height)

		left, right := 0, size-1

		maxWidth := size - 1

		area := 0

		for width := maxWidth; width > 0; width-- {
			if height[left] < height[right] {
				area = Max(area, width*height[left])
				left += 1
			} else {
				area = Max(area, width*height[right])
				right -= 1
			}
		}
		return area
	}
	return 0
}

func Max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}
