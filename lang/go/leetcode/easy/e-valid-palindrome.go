package leetcode

/**
验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

https://leetcode-cn.com/problems/valid-palindrome/
*/
//与有效括号一样,栈都不用，直接双指针
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	var start = 0
	var end = len(s) - 1

	isLetter := func(a byte) bool { return (a >= 'a' && a <= 'z') || (a >= 'A' || a <= 'Z') }
	eqLetter := func(a, b byte) bool {
		return a == b || (isLetter(a) && isLetter(b) && (s[start]-32 == s[end] || s[end]-32 == s[start]))
	}
	for {
		//比起这种写法s[start] !in 'a'..'z' && s[start] !in 'A'..'Z' && s[start] !in '0'..'9'，效率高
		for !(isLetter(s[start]) || (s[start] >= '0' && s[start] <= '9')) {
			if start == end {
				return true
			}
			start++
		}
		for !(isLetter(s[end]) || (s[end] >= '0' && s[end] <= '9')) {
			end--
		}
		if start >= end {
			return true
		}
		if eqLetter(s[start], s[end]) {
			start++
			end--
		}
		return false
	}
}
