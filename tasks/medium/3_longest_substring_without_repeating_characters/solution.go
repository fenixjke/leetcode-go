package task

import (
	"leetcode-go/common"
	"strings"
)

type S struct{}

var option string = "leetcode-2"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"straight", "leetcode", "leetcode-with-map", "leetcode-2"},
		TaskName: "Longest Substring Without Repeating Characters",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input) {
		_ = lengthOfLongestSubstring(input[i])
	}
}

func lengthOfLongestSubstring(s string) int {
	switch option {
	case "straight":
		var max int
		var sub string
		var i int
		for i < len(s) {
			if i >= len(s) {
				break
			}
			sb := s[i : i+1]
			i += 1
			lt := strings.LastIndex(sub, sb)
			sub += sb
			if lt < 0 {
				if len(sub) > max {
					max = len(sub)
				}
			} else {
				sub = sub[lt+1:]
			}
		}
		return max
	case "leetcode":
		dict := [128]bool{}
		length, max := 0, 0
		for i, j := 0, 0; i < len(s); i++ {
			index := s[i]
			if dict[index] {
				for ; dict[index]; j++ {
					length--
					dict[s[j]] = false
				}
			}

			dict[index] = true
			length++
			if length > max {
				max = length
			}
		}
		return max
	case "leetcode-with-map":
		curr := map[byte]bool{}
		m := 0
		for i, j, l := 0, 0, 0; i < len(s); i++ {
			sb := s[i]
			for ; curr[sb]; j++ {
				curr[s[j]] = false
				l--
			}

			curr[sb] = true
			l++
			if m < l {
				m = l
			}
		}
		return m
	case "leetcode-2":
		start := 0
		longest := 0
		used := map[byte]int{}

		for i := 0; i < len(s); i++ {
			c := s[i]

			if _, ok := used[c]; ok && used[c] >= start {
				start = used[c] + 1
			}

			longest = max(longest, i-start+1)
			used[c] = i
		}

		return longest
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
