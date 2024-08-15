package task

import (
	"leetcode-go/common"
	"strconv"
)

type S struct{}

var option string = "my-devide"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my-to-str", "my-devide"},
		TaskName: "Palindrome Number",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = isPalindrome(input1[i])
	}
}

func isPalindrome(x int) bool {
	switch option {
	case "my-to-str":
		str := strconv.Itoa(x)
		return str == revertStr(str)
	case "my-devide":
		if x < 0 {
			return false
		}
		temp := x
		var reverted int
		for temp != 0 {
			reverted = reverted*10 + temp%10
			temp /= 10
		}
		return x == reverted
	}

	return false
}

func revertStr(s string) string {
	b := []byte(s)
	br := make([]byte, len(b))
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		br[j] = b[i]
	}
	return string(br)
}
