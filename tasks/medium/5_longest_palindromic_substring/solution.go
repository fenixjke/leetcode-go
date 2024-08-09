package task

import (
	"leetcode-go/common"
)

type S struct{}

var option string = "straight"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"straight", "leetcode", "leetcode-1"},
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

func lengthOfLongestSubstring(s string) string {
	switch option {
	case "straight":
		res := ""
		max, last := 0, 0
		i := 1
		for len(s)-last > max {
			i = len(s)
			for max < i-last {
				subs := s[last:i]
				b := []byte(subs)
				poly := true
				for o, p := 0, len(b)-1; o < p; o, p = o+1, p-1 {
					if b[o] != b[p] {
						poly = false
						break
					}
				}
				if poly {
					str := string(b)
					if len(str) > max {
						res = str
						max = len(str)
					}
				}
				i--
			}
			last++
		}
		return res
	case "leetcode":
		sPrime := "#"
		for _, c := range s {
			sPrime = sPrime + string(c) + "#"
		}

		n := len(sPrime)
		palindromeRadii := make([]int, n)
		center := 0
		radius := 0

		for i := 0; i < n; i++ {
			mirror := 2*center - i

			if i < radius {
				palindromeRadii[i] = min(radius-i, palindromeRadii[mirror])
			}

			for i+1+palindromeRadii[i] < n && i-1-palindromeRadii[i] >= 0 &&
				sPrime[i+1+palindromeRadii[i]] == sPrime[i-1-palindromeRadii[i]] {
				palindromeRadii[i]++
			}

			if i+palindromeRadii[i] > radius {
				center = i
				radius = i + palindromeRadii[i]
			}
		}

		maxLength := 0
		centerIndex := 0
		for i := 0; i < n; i++ {
			if palindromeRadii[i] > maxLength {
				maxLength = palindromeRadii[i]
				centerIndex = i
			}
		}

		startIndex := (centerIndex - maxLength) / 2
		longestPalindrome := s[startIndex : startIndex+maxLength]

		return longestPalindrome
	case "leetcode-1":
		ll := len(s)
		if ll == 0 {
			return ""
		}

		var l, r, pl, pr int
		for r < ll {
			// gobble up dup chars
			for r+1 < ll && s[l] == s[r+1] {
				r++
			}
			// find size of this palindrome
			for l-1 >= 0 && r+1 < ll && s[l-1] == s[r+1] {
				l--
				r++
			}
			if r-l > pr-pl {
				pl, pr = l, r
			}
			// reset to next mid point
			l = (l+r)/2 + 1
			r = l
		}
		return s[pl : pr+1]
	}

	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
