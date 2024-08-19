package task

var input1 = []string{
	"aa",
	"aa",
	"ab",
	"aab",
	"aabccmcacccacg",
	"ab",
	"a",
	"b",
}

var input2 = []string{
	"a",
	"a*",
	".*",
	"c*a*b",
	"ac*.*m.*a*acg",
	".*..",
	"ab*",
	"aaa.",
}

var expects = []bool{
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
}
