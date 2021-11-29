package kmp

func getNext(pattern string) []int {
	if len(pattern) == 0 {
		return nil
	}

	next := make([]int, len(pattern))
	next[0] = 0
	// k为匹配的前后缀最大公共子串的长度计数，与前缀的索引也能对应
	k := 0
	// i表示模式串pattern的左子串的最右边的字符索引，遍历处理所有模式串的左子串
	i := 1
	for i < len(pattern) {
		// 这时i>0，至少是两个字符，例如"aa"，前后缀最大公共子串为"a"
		// 两个索引对应的字符相同时，同时递增这两索引，再比较它们后面的字符
		if pattern[k] == pattern[i] {
			// k作为索引递增时，长度也相应递增
			k += 1
			next[i] = k
			i += 1
		} else {
			// 当两个索引对应字符不等时，例如"ab"，"abc"，"abcd"
			if k > 0 {
				// 已有部分匹配，则把k重新回朔到开头
				k = 0
			} else {
				// 没有任何匹配时，继续让i递增，k不变
				i += 1
			}
		}
	}

	return next
}

// kmp算法，返回匹配到的字符串开始索引列表
func Kmp(str, pat string) []int {
	if len(str) == 0 || len(pat) == 0 || len(str) < len(pat) {
		return nil
	}

	next := getNext(pat)

	var starts []int
	i := 0
	j := 0
	for i < len(str) {
		if str[i] == pat[j] {
			i += 1
			j += 1
		} else {
			// 模式串"ab"，主串"aaaaaa"
			// 在之前有部分匹配的情况下，把模式串的索引挪到next值对应的索引上
			if j > 0 {
				j = next[j-1]
			} else {
				i += 1
			}
		}
		if j == len(pat) {
			// 已经完全匹配
			starts = append(starts, i-len(pat))
			// 重置j为next值
			j = next[j-1]
		}
	}
	return starts
}
