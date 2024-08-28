package task

import (
	"leetcode-go/common"
)

type S struct{}

var option string = "my"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my", "leetcode"},
		TaskName: "Roman to Integer",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = romanToInt(input1[i])
	}
}

var symbols = []string{
	"I", "V", "X", "L", "C", "D", "M",
}

var numbersTable = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	switch option {
	case "my":
		r := 0
		for i := 0; i < len(s); i++ {
			c := s[i]
			var n byte
			if len(s) > i+1 {
				n = s[i+1]
			}
			if n != 0 {
				if v, ok := numbersTable[string(c)+string(n)]; ok {
					r += v
					i++
				} else {
					r += numbersTable[string(c)]
				}
			} else {
				r += numbersTable[string(c)]
			}
		}
		return r
	case "leetcode":
		var v, lv, cv int
		h := map[uint8]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}

		for i := len(s) - 1; i >= 0; i-- {
			cv = h[s[i]]
			if cv < lv {
				v -= cv
			} else {
				v += cv
			}
			lv = cv
		}

		return v
	}
	return 0
}
