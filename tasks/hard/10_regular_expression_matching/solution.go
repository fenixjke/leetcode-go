package task

import (
	"leetcode-go/common"
)

type S struct{}

var option string = "my-straigh"

func New() common.TaskRunner {
	return &common.Info{
		Options:  []string{"my-straigh", "leetcode", "leetcode-2"},
		TaskName: "Regular Expression Matching",
		Runner:   &S{},
	}
}

func (s *S) Run(o string) {
	option = o
	for i := range len(input1) {
		_ = isMatch(input1[i], input2[i])
	}
}

var si, pi int

func isMatch(s string, p string) bool {
	switch option {
	case "my-straigh":
		return matchRemaining(s, p, si, pi)
	case "leetcode":
		slen, plen := len(s), len(p)
		var dp [][]bool
		var t []bool
		for i := 0; i <= slen; i++ {
			t = make([]bool, plen+1)
			dp = append(dp, t)
		}

		for i := 0; i <= slen; i++ {
			for j := 0; j <= plen; j++ {
				if i == 0 && j == 0 {
					dp[i][j] = true
					continue
				} else if i == 0 {
					dp[i][j] = ((j-1)%2 == 1 && p[j-1] == '*' && dp[i][j-2])
					continue
				} else if j == 0 {
					dp[i][j] = false
					continue
				}

				if p[j-1] != '*' {
					dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '.') && dp[i-1][j-1]
				} else {
					if p[j-2] == '.' || p[j-2] == s[i-1] {
						dp[i][j] = dp[i-1][j]
					}
					if dp[i][j-2] == true {
						dp[i][j] = true
					}
				}
			}
		}
		return dp[slen][plen]
	case "leetcode-2":
		// 1.
		table := make([][]bool, len(s)+1)
		for i := 0; i < len(table); i++ {
			table[i] = make([]bool, len(p)+1)
		}
		table[0][0] = true

		// 2.
		for j := 2; j < len(p)+1; j++ {
			if p[j-1] == '*' {
				table[0][j] = table[0][j-2]
			}
		}

		// 3.
		for i := 1; i < len(s)+1; i++ {
			for j := 1; j < len(p)+1; j++ {
				if s[i-1] == p[j-1] || p[j-1] == '.' {
					// 4.
					table[i][j] = table[i-1][j-1]
				} else if p[j-1] == '*' {
					// 5.
					empty := table[i][j-2]
					nonempty := (s[i-1] == p[j-2] || p[j-2] == '.') && table[i-1][j]
					table[i][j] = empty || nonempty
				}
			}
		}

		// 6.
		return table[len(s)][len(p)]
	}
	return false
}

func matchRemaining(s, p string, si, pi int) bool {
	if len(s) == si && len(p) == pi {
		return true
	}

	originalsi := si
	originalpi := pi

	var token byte
	token, pi = getNextToken(p, pi)
	zom := pi-originalpi > 1

	if (len(s) == si && !zom) || len(p) == originalpi {
		return false
	}

	var ok, relax bool

	for {
		if !ok && zom {
			ok = true
		}

		if len(s) == si {
			break
		}

		if token == '.' || token == s[si] {
			si++
			ok = true
			if zom {
				relax = true
			} else {
				break
			}
		} else {
			break
		}
	}

	if ok {
		ok = matchRemaining(s, p, si, pi)
	}

	for ; !ok; ok = matchRemaining(s, p, si, pi) {
		if relax && (originalsi < si) {
			si--
		} else {
			break
		}
	}

	return ok
}

func getNextToken(p string, pi int) (byte, int) {
	if len(p) == pi {
		return 0, pi
	}

	t := p[pi]
	pi++

	zom := (len(p) > pi) && p[pi] == '*'

	if zom {
		pi++
	}

	return t, pi
}
