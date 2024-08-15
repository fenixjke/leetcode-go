package task

import (
	"leetcode-go/common"
	"math"
	"strconv"
)

type S struct{}

var option string = "leetcode"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my", "leetcode"},
		TaskName: "String to Integer (atoi)",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = myAtoi(input1[i])
	}
}

var step1Chars = []byte{' '}
var step2Chars = []byte{'+', '-'}
var step3Chars = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var digits = map[byte]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}

func myAtoi(s string) int {
	switch option {
	case "my":
		var res, temp int
		var mult int = 1
		step := 1
		for i := 0; i < len(s); i++ {
			b := s[i]
			if (step == 2) || (step < 3 && contains(step3Chars, b)) {
				step = 3
			} else if step < 2 && contains(step2Chars, b) {
				step = 2
			} else if step == 1 && !contains(step1Chars, b) {
				break
			}
			switch step {
			case 1:
				continue
			case 2:
				if b == step2Chars[1] {
					mult = -1
				}
			case 3:
				if contains(step3Chars, b) {
					c, _ := strconv.Atoi(string(b))
					temp = temp*10 + c
					if res < temp {
						res = temp
					} else if res > temp {
						res = math.MaxInt32 + 2
						break
					}
				} else {
					i = len(s)
				}
			}
		}
		res = res * mult
		if res < math.MinInt32 {
			res = math.MinInt32
		} else if res > math.MaxInt32 {
			res = math.MaxInt32
		}
		return res
	case "leetcode":
		res, sign, len, idx := 0, 1, len(s), 0

		// Skip leading spaces
		for idx < len && (s[idx] == ' ' || s[idx] == '\t') {
			idx++
		}

		if idx == len {
			return 0
		}

		// +/- Sign
		if s[idx] == '+' {
			sign = 1
			idx++
		} else if s[idx] == '-' {
			sign = -1
			idx++
		}

		// Digits: 0x30 = '0', 0x31 = '1', ... 0x39 = '9'
		for idx < len && s[idx] >= 0x30 && s[idx] <= 0x39 {
			res = res*10 + digits[s[idx]]
			if sign*res > math.MaxInt32 {
				return math.MaxInt32
			}

			if sign*res < math.MinInt32 {
				return math.MinInt32
			}

			idx++
		}

		return res * sign
	}
	return 0
}

func contains(arr []byte, b byte) bool {
	for i := range arr {
		if arr[i] == b {
			return true
		}
	}
	return false
}
