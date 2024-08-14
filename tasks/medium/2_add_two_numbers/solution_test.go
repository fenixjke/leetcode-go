package task

import (
	"leetcode-go/common"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		l1   *common.ListNode
		l2   *common.ListNode
		want *common.ListNode
	}{
		{l1s[0], l2s[0], expects[0]},
		{l1s[1], l2s[1], expects[1]},
		{l1s[2], l2s[2], expects[2]},
	}
	for _, c := range cases {
		got := addTwoNumbers(c.l1, c.l2)
		if got == nil {
			t.Errorf("Result is nil")
		}
		for c.want != nil || got != nil {
			if c.want == nil {
				t.Errorf("Not expected value in result: %v", got)
				t.FailNow()
			} else if got == nil {
				t.Errorf("Result has nil but want: %v", c.want)
				t.FailNow()
			}
			if c.want.Val != got.Val {
				t.Errorf("L1 %v, L2 %v, expected result %v, but result is %v", c.l1, c.l2, c.want, got)
				t.FailNow()
			}
			c.want = c.want.Next
			got = got.Next
		}
	}
}
