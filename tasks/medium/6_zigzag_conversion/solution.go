package task

import (
	"leetcode-go/common"
	"math"
	"strings"
)

type S struct{}

var option string = "leetcode"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"straight", "leetcode"},
		TaskName: "Zigzag Conversion",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = convert(input1[i], input2[i])
	}
}

func convert(s string, numRows int) string {
	switch option {
	case "straight":
		var arrLength int
		if numRows == 1 {
			arrLength = len(s)
		} else {
			arrLength = calcArrLength(s, numRows)
		}
		arrs := make([][]byte, numRows)
		for i := range arrs {
			arrs[i] = make([]byte, arrLength)
		}

		var down bool
		for i, a, b := 0, 0, 0; i < len(s); i += 1 {
			if a+1 == numRows {
				down = false
			} else if a == 0 {
				down = true
			}
			arrs[a][b] = s[i]
			if down {
				a++
			} else {
				if a > 0 {
					a--
				}
				b++
			}
		}
		b := []byte{}
		for a1 := 0; a1 < len(arrs); a1++ {
			for a2 := 0; a2 < len(arrs[a1]); a2++ {
				if arrs[a1][a2] != 0 {
					b = append(b, arrs[a1][a2])
				}
			}
		}
		return string(b)
	case "leetcode":
		if numRows == 1 {
			return s
		}

		var b strings.Builder
		b.Grow(len(s))
		step := numRows*2 - 2

		for row := 0; row < numRows; row++ {
			nextStep := 0
			if row == 0 || row == numRows-1 {
				nextStep = step
			} else {
				nextStep = row * 2
			}
			for pos := row; pos < len(s); pos += nextStep {
				b.WriteByte(s[pos])
				if row == 0 || row == numRows-1 {
					nextStep = step
				} else {
					nextStep = step - nextStep
				}
			}
		}

		return b.String()
	}
	return ""
}

func calcArrLength(s string, n int) int {
	fc := len(s) / (2*n - 2)
	rem := len(s) - (fc * (2*n - 2))
	reml := int(math.Min(float64(rem), math.Max(float64(rem-n), 0.0)+1))
	ml := fc + fc*(n-2)
	return reml + ml
}
