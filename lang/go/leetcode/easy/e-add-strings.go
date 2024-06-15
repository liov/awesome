package leetcode

import (
	"bytes"
	"github.com/hopeio/cherry/utils/slices"
	stringsi "github.com/hopeio/cherry/utils/strings"
)

/*
*
415. 字符串相加

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	var short, long = num1, num2
	if len(num1) > len(num2) {
		short, long = num2, num1
	}
	m, n := len(short), len(long)
	var carry = byte(0)
	var ret = bytes.Buffer{}
	ret.Grow(n + 1)
	for i := range n {
		sum := (long[n-i] - '0') + carry
		if i < m {
			sum += short[m-i] - '0'
		}

		if sum >= 10 {
			carry = 1
			ret.WriteByte('0' + sum - 10)
		} else {
			carry = 0
			ret.WriteByte('0' + sum)
		}
	}
	if carry == 1 {
		ret.WriteByte('1')
	}
	return stringsi.ToString(slices.Reverse(ret.Bytes()))
}
