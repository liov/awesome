package leetcode

import (
	"bytes"
	"github.com/hopeio/cherry/utils/slices"
	stringsi "github.com/hopeio/cherry/utils/strings"
)

/*
*
67. 二进制求和

给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addBinary(a string, b string) string {
	if a == "0" {
		return b
	}

	if b == "0" {
		return a
	}
	maxLen := len(b)
	if len(a) > len(b) {
		maxLen = len(a)
	}
	var carry = false
	var oneCount = byte(0) //1出现的次数，偶数该位置0且进一位
	var ret = bytes.Buffer{}
	ret.Grow(maxLen + 1)
	var aIdx, bIdx int
	for i := range maxLen {
		aIdx = len(a) - 1 - i
		bIdx = len(b) - 1 - i
		if aIdx >= 0 && a[aIdx] == '1' {
			oneCount++
		}
		if bIdx >= 0 && b[bIdx] == '1' {
			oneCount++
		}
		if carry {
			oneCount++
		}
		ret.WriteByte((oneCount & 1) + '0')
		carry = oneCount > 1 //两个或三个1进1位
		oneCount = 0
	}
	if carry {
		ret.WriteByte('1')
	}
	return stringsi.ToString(slices.Reverse(ret.Bytes()))
}
