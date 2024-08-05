package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var r []int
	for l != nil {
		r = append(r, l.Val)
		l = l.Next
	}
	return fmt.Sprintf("%v", r)
}
