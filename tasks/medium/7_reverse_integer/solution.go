package task

import (
	"leetcode-go/common"
	"math"
	"strconv"
	"strings"
)

type S struct{}

var option string = "my-to-string"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my-devide-approach", "my-to-string", "leetcode"},
		TaskName: "Reverse Integer",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = reverse(input1[i])
	}
}

func reverse(x int) int {
	switch option {
	case "my-devide-approach":
		var res, fract int
		for x != 0 {
			x, fract = getIntAndFractPart(x)
			res = res*10 + fract
		}

		if res < math.MinInt32 || res > math.MaxInt32 {
			res = 0
		}

		return res
	case "my-to-string":
		str := strconv.Itoa(x)
		strs := strings.Split(str, "")
		strr := ""
		mult := 1
		for i := len(strs) - 1; i >= 0; i-- {
			if strs[i] == "-" {
				mult = -1
			} else {
				strr += strs[i]
			}
		}
		a, _ := strconv.Atoi(strr)

		res := a * mult

		if res < math.MinInt32 || res > math.MaxInt32 {
			res = 0
		}

		return res
	case "leetcode":
		var result int
		for x != 0 {
			result = result*10 + x%10
			if result > 2147483647 || result < -2147483648 {
				return 0
			}
			x /= 10
		}
		return result
	}
	return 0
}

func getIntAndFractPart(x int) (int, int) {
	quotient := float64(x) / 10.0
	return int(quotient), x - int(quotient)*10
}
