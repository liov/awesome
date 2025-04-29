package leetcode

/**
718. 最长重复子数组

给两个整数数组A和B，返回两个数组中公共的、长度最长的子数组的长度。

示例 1:

输入:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
输出: 3
解释:
长度最长的公共子数组是 [3, 2, 1]。
说明:

1 <= len(A), len(B) <= 1000
0 <= A[i], B[i] < 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

const mod int = 1000000009
const base = 113

func findLength(A []int, B []int) int {
	var left = 1
	var right = min(len(A), len(B)) + 1
	for left < right {
		var mid = (left + right) >> 1
		if check(A, B, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

// 字符串匹配，Rabin-Karp
// 其中，base代表的是S的定义域大小，比如说如果S全是英文字母，那么的值为26，因为英文字母就只有26个。然后这个函数是一个映射函数，映射S的定义域中的每一个字符到数字的函数。
// 如果采用O(m)的算法计算长度为m的字符串子串的哈希值的话，那复杂度还是O(nm)。这是就要使用一个滚动哈希的优化技巧。
// 选取两个合适的互素常数b和h(l<b<h)，假设字符串C=c1c2...cm，定义哈希函数：
// H(C)=(c1*b^(m-1)+c2*b^(m-2))+c3*b^(m-3)+...+cm*b^0)%h
// 其中b是基数，相当于把字符串看作b进制数。这样，字符串S=s1s2...sn从位置k+1开始长度为m的字符串子串S[k+1...k+m]的哈希值，就可以利用从位置k开始的字符串子串S[k...k+m-1]的哈希值直接进行计算：
// H(S[k+1...k+m])=(H(s[k:k+m-1])*b-sk*b^m+s(k+m))%h
// 二分+hash，逐步缩小
// hash(S)= i=0∑∣S∣−1 base^∣S∣−(i+1)×S[i]
// hash(S[1:len+1])=(hash(S[0:len])−base^(len−1)×S[0])×base+S[len+1]
func check(A []int, B []int, l int) bool {
	var hashA = 0
	//0到len-1的hash
	for i := range l {
		hashA = (hashA*base + A[i]) % mod
	}
	var bucketA = make(map[int]struct{})
	bucketA[hashA] = struct{}{}
	var mult = qPow(base, l)
	//1到len，2到len+1...
	for i := l; i < len(A); i++ {
		hashA = (hashA*base - A[i-l]*mult + A[i]) % mod
		if hashA < 0 {
			hashA += mod
		}
		bucketA[hashA] = struct{}{}
	}
	var hashB = 0
	for i := range l {
		hashB = (hashB*base + B[i]) % mod
	}
	if _, ok := bucketA[hashB]; ok {
		return true
	}
	for i := l; i < len(B); i++ {
		hashB = (hashB*base - B[i-l]*mult + B[i]) % mod
		if hashB < 0 {
			hashA += mod
		}
		if _, ok := bucketA[hashB]; ok {
			return true
		}
	}
	return false
}

// 使用快速幂计算 x^n % mod 的值
func qPow(x int, n int) int {
	var ret = 1
	for n != 0 {
		if n&1 != 0 {
			ret = ret * x % mod
		}
		x = x * x % mod
		n = n >> 1
	}
	return ret
}
