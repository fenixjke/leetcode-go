package twosum

import "leetcode-go/common"

type S struct{}

var option string = "straight"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"straight", "leetcode"},
		TaskName: "Add Two Numbers",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(l1s) {
		_ = addTwoNumbers(l1s[i], l2s[i])
	}
}

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	switch option {
	case "straight":
		return process(l1, l2, 0)
	case "leetcode":
		carry, dummy := 0, new(common.ListNode)
		for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
			if l1 != nil {
				carry += l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				carry += l2.Val
				l2 = l2.Next
			}
			node.Next = &common.ListNode{carry % 10, nil}
			carry /= 10
		}
		return dummy.Next
	}
	return nil
}

func process(l1 *common.ListNode, l2 *common.ListNode, addition int) *common.ListNode {
	l := &common.ListNode{}
	var r = addition
	if l1 != nil {
		r += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		r += l2.Val
		l2 = l2.Next
	}
	a := r / 10
	c := r % 10
	l.Val = c
	if l1 != nil || l2 != nil || a != 0 {
		l.Next = process(l1, l2, a)
	}
	return l
}
