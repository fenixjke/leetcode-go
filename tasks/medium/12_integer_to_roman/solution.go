package task

import (
	"leetcode-go/common"
)

type S struct{}

var option string = "my"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my", "leetcode"},
		TaskName: "Integer to Roman",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = intToRoman(input1[i])
	}
}

var numbers = []int{
	1, 5, 10, 50, 100, 500, 1000,
}

var symbolTable = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	switch option {
	case "my":
		res := ""
		mult := 1
		for num > 0 {
			rem := (num % 10) * mult
			temp := ""
			for rem > 0 {
				if rem != 4*mult && rem != 9*mult {
					max, symb := getMax(rem)
					temp += symb
					rem -= max
				} else {
					temp = getSpec(rem, mult)
					rem = 0
				}
			}
			res = temp + res
			num = num / 10
			mult *= 10
		}
		return res
	case "leetcode":
		return getRoman(num, "")
	}
	return ""
}

func getMax(num int) (int, string) {
	max := 1
	for _, v := range numbers {
		if v <= num && v >= max {
			max = v
		} else {
			break
		}
	}
	return max, symbolTable[max]
}

func getSpec(num, mult int) string {
	next := 1
	for _, v := range numbers {
		if v < num {
			continue
		} else {
			next = v
			break
		}
	}
	prev := 1 * mult
	return symbolTable[prev] + symbolTable[next]
}

func getRoman(num int, roman string) string {
	if num == 0 {
		return roman
	}

	if num-1000 >= 0 {
		return getRoman(num-1000, roman+"M")
	}

	if num-900 >= 0 {
		return getRoman(num-900, roman+"CM")
	}

	if num-500 >= 0 {
		return getRoman(num-500, roman+"D")
	}

	if num-400 >= 0 {
		return getRoman(num-400, roman+"CD")
	}

	if num-100 >= 0 {
		return getRoman(num-100, roman+"C")
	}

	if num-90 >= 0 {
		return getRoman(num-90, roman+"XC")
	}

	if num-50 >= 0 {
		return getRoman(num-50, roman+"L")
	}

	if num-40 >= 0 {
		return getRoman(num-40, roman+"XL")
	}

	if num-10 >= 0 {
		return getRoman(num-10, roman+"X")
	}

	if num-9 >= 0 {
		return getRoman(num-9, roman+"IX")
	}

	if num-5 >= 0 {
		return getRoman(num-5, roman+"V")
	}

	if num-4 >= 0 {
		return getRoman(num-4, roman+"IV")
	}

	if num-1 >= 0 {
		return getRoman(num-1, roman+"I")
	}

	return roman
}
