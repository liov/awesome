package leetcode

import "math"

/*
29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2

提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。

https://leetcode-cn.com/problems/divide-two-integers/
*/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	retGtZero := true
	if dividend < 0 {
		dividend = -dividend
		retGtZero = !retGtZero
	}
	if divisor < 0 {
		divisor = -divisor
		retGtZero = !retGtZero
	}
	if dividend < divisor {
		return 0
	}
	if divisor == 1 {
		if retGtZero {
			return dividend
		}
		return -dividend
	}
	if dividend == divisor {
		if retGtZero {
			return 1
		}
		return -1
	}
	var ret int
	for i := 31; i >= 0; i-- {
		if divisor<<i <= dividend {
			dividend -= divisor << i
			ret += 1 << i
		}
	}
	if retGtZero {
		return ret
	}
	return -ret
}
