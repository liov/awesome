package leetcode

/**
28. 实现 strStr()
实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。



说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。



示例 1：

输入：haystack = "hello", needle = "ll"
输出：2
示例 2：

输入：haystack = "aaaaa", needle = "bba"
输出：-1
示例 3：

输入：haystack = "", needle = ""
输出：0


提示：

0 <= haystack.length, needle.length <= 5 * 104
haystack 和 needle 仅由小写英文字符组成
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	return rabinKarp(haystack, needle)
}

func KMP(raw string, pat string) int {
	next := buildNext(pat)
	n, m := len(raw), len(pat)
	var i = 0

	var j = 0
	for j < m && i < n {
		if 0 > j || raw[i] == pat[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		return i - j
	}
	return -1
}

func buildNext(pat string) []int {
	m := len(pat)
	var j = 0
	next := make([]int, m)
	var t = -1
	next[0] = t
	for j < m-1 {
		if 0 > t || pat[j] == pat[t] {
			j++
			t++
			next[j] = next[t]
			if pat[j] != pat[t] {
				next[j] = t
			}
		} else {
			t = next[t]
		}
	}
	return next
}

func Sunday(raw string, pat string) int {
	n, m := len(raw), len(pat)
	var i = 0
	var j = 0
	offset := make(map[byte]int)
	for c := range pat {
		offset[pat[c]] = m - c
	}
	for i <= n-m {
		j = 0
		for raw[i+j] == pat[j] {
			j += 1
			if j == m {
				return i
			}
		}
		if i+m == n {
			return -1
		}
		if _, ok := offset[raw[i+m]]; ok {
			i += offset[raw[i+m]]
		} else {
			i += m + 1
		}
	}
	return -1
}

func preBmBc(x string, bmBc []int) {
	// 计算字符串中每个字符距离字符串尾部的长度
	for i := range x {
		bmBc[x[i]] = len(x) - i - 1
	}
}

// 计算以i为边界，与模式串后缀匹配的最大长度（区间的概念）
func suffix(x string, suff []int) {
	n := len(x)
	var q int
	for i := n - 2; i >= 0; i-- {
		q = i
		for q >= 0 && x[q] == x[n-1-i+q] {
			q--
		}
		suff[i] = i - q
	}
}

// 好后缀算法的预处理
/*
 有三种情况
 1.模式串中有子串匹配上好后缀
 2.模式串中没有子串匹配上好后缀，但找到一个最大前缀
 3.模式串中没有子串匹配上好后缀，但找不到一个最大前缀


 3种情况获得的bmGs[i]值比较

 3 > 2 > 1

 为了保证其值越来越小

 所以按顺序处理3->2->1情况
*/
func preBmGs(s string, bmGs []int) {

	n := len(s)
	suff := make([]int, n)

	suffix(s, suff)

	//全部初始为自己的长度，处理第三种情况
	/*
	   for (i in s.indices) {
	     bmGs[i] = len
	   }
	*/

	// 处理第二种情况
	for i := n - 1; i >= 0; i-- {
		if suff[i] == i+1 { // 找到合适位置
			for j := range s {
				if bmGs[j] == n {
					bmGs[j] = n - 1 - i // 保证每个位置至多只能被修改一次
				}
			}
		}
	}

	// 处理第一种情况，顺序是从前到后
	for i := range s {
		bmGs[n-1-suff[i]] = n - 1 - i
	}

}

func BM(raw string, pat string) int {

	n, m := len(raw), len(pat)

	bmGs := make([]int, m)
	for i := range bmGs {
		bmGs[i] = m
	}
	// 全部更新为自己的长度
	bmBc := make([]int, 256)
	for i := range bmBc {
		bmBc[i] = m
	}
	// 处理好后缀算法
	preBmGs(pat, bmGs)
	// 处理坏字符算法
	preBmBc(pat, bmBc)

	var j = 0

	for j <= n-m {
		// 模式串向左边移动
		var i = m - 1
		for i >= 0 && pat[i] == raw[i+j] {
			i--
		}
		// 给定字符串向右边移动
		if i < 0 {
			return j // 移动到模式串的下一个位置
		} else {
			// 取移动位数的最大值向右移动，前者好后缀，向右移动，后者坏字符，向左移动
			j += max(bmGs[i], bmBc[raw[i+j]]-m+1+i)
		}
	}
	return -1
}

//jvm没有无符号整型
//go的子串算法

func RabinKarp(s string, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h = uint(0)
	for i := range substr {
		h = h*16777619 + uint(s[i])
	}
	if h == hashss && s[0:n] == substr {
		return 0
	}

	for i := n; i <= len(s); i++ {
		h = h*16777619 + uint(s[i]) - pow*uint(s[i-n])
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

func hashStr(sep string) (uint, uint) {
	var hash = uint(0)
	for i := range sep {
		hash = hash*16777619 + uint(sep[i])
	}
	var pow = uint(1)
	var sq = uint(16777619)
	var i = len(sep)
	for i > 0 {
		i = i >> 1
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func rabinKarp(str string, pattern string) int {
	n, m := len(str), len(pattern)

	//哈希时需要用到进制计算，这里只涉及26个字母所以使用26进制
	d := uint(26)
	//防止hash之后的值超出int范围，对最后的hash值取模
	//q取随机素数，满足q*d < INT_MAX即可
	q := uint(16777619)

	//str子串的hash值
	var strCode = uint(0)
	//pattern的hash值
	var patternCode = uint(0)
	//d的size2-1次幂，hash计算时，公式中会用到
	var h = uint(1)

	//计算sCode、pCode、h
	for i := range m {
		patternCode = (d*patternCode + uint(pattern[i]-'a')) % q
		//计算str第一个子串的hash
		strCode = (d*strCode + uint(str[i]-'a')) % q
		h = (h * d) % q
	}
	//最大需要匹配的次数
	//字符串开始匹配，对patternCode和strCode开始比较，并更新strCode的值
	for i := range n - m + 1 {
		if strCode == patternCode && ensureMatching(i, str, pattern) {
			return i
		}
		if i == n-m {
			break
		}
		//更新strCode的值，即计算str[i+1,i+m-1]子串的hashCode
		strCode = (strCode*d - h*uint(str[i]-'a') + uint(str[i+m]-'a')) % q
	}
	return -1
}

/**
 * hash值一样并不能完全确保字符串一致，所以还需要进一步确认
 * @param i hash值相同时字符串比对的位置
 * @param pattern 模式串
 * @return
 */
func ensureMatching(i int, str string, pattern string) bool {
	strSub := str[i : i+len(pattern)]
	return strSub == pattern
}
