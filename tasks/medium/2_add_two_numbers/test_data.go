package task

import "leetcode-go/common"

var l1s = []*common.ListNode{
	{
		Val: 2,
		Next: &common.ListNode{
			Val: 4,
			Next: &common.ListNode{
				Val: 3,
			},
		},
	},
	{
		Val: 0,
	},
	{
		Val: 9,
		Next: &common.ListNode{
			Val: 9,
			Next: &common.ListNode{
				Val: 9,
				Next: &common.ListNode{
					Val: 9,
					Next: &common.ListNode{
						Val: 9,
						Next: &common.ListNode{
							Val: 9,
							Next: &common.ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	},
}

var l2s = []*common.ListNode{
	{
		Val: 5,
		Next: &common.ListNode{
			Val: 6,
			Next: &common.ListNode{
				Val: 4,
			},
		},
	},
	{
		Val: 0,
	},
	{
		Val: 9,
		Next: &common.ListNode{
			Val: 9,
			Next: &common.ListNode{
				Val: 9,
				Next: &common.ListNode{
					Val: 9,
				},
			},
		},
	},
}

var expects = []*common.ListNode{
	{
		Val: 7,
		Next: &common.ListNode{
			Val: 0,
			Next: &common.ListNode{
				Val: 8,
			},
		},
	},
	{
		Val: 0,
	},
	{
		Val: 8,
		Next: &common.ListNode{
			Val: 9,
			Next: &common.ListNode{
				Val: 9,
				Next: &common.ListNode{
					Val: 9,
					Next: &common.ListNode{
						Val: 0,
						Next: &common.ListNode{
							Val: 0,
							Next: &common.ListNode{
								Val: 0,
								Next: &common.ListNode{
									Val: 1,
								},
							},
						},
					},
				},
			},
		},
	},
}
