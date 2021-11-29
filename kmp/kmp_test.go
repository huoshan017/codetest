package kmp

import (
	"testing"
)

// 通过这个函数来测试Kmp函数
func TestGetNext(t *testing.T) {
	pat := "aacdeaacdabcda"
	next := getNext(pat)
	t.Logf("pattern %v next array is: %v", pat, next)
	str := "baacdeaacdabcdaaacdeaacdabcd"
	starts := Kmp(str, pat)
	t.Logf("string %v match pattern %v starts: %v", str, pat, starts)

	pat = "bb"
	t.Logf("pattern %v next array is: %v", pat, getNext(pat))
	starts = Kmp(str, pat)
	t.Logf("string %v match pattern %v starts: %v", str, pat, starts)

	pat = "abbbaa"
	str = "aabaaabaaaababaabaaaabbbbbbaaaaaccddeddedacaaabbbbaabbbabb"
	t.Logf("string %v match pattern %v starts: %v", str, pat, Kmp(str, pat))

	pat = "rw.rm.Unlock()"
	t.Logf("string %v match pattern %v starts: %v", test_str, pat, Kmp(test_str, pat))
}
