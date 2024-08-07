package task

import "leetcode-go/common"

type S struct{}

var option string = "straight"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"straight", "complement"},
		TaskName: "Two Sum",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(InputArrays) {
		_ = twoSum(InputArrays[i], Targets[i])
	}
}

func twoSum(nums []int, target int) []int {
	switch option {
	case "straight":
		for i, val1 := range nums {
			for j, val2 := range nums[i+1:] {
				if val1+val2 == target {
					return []int{i, i + j + 1}
				}
			}
		}
	case "complement":
		seen := make(map[int]int)
		for k1, v := range nums {
			complement := target - v
			if k2, ok := seen[complement]; ok {
				return []int{k2, k1}
			}

			seen[v] = k1
		}
	}

	return nil
}
