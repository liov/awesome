package leetcode

/**
154. 寻找旋转排序数组中的最小值 II

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

注意数组中可能存在重复的元素。

示例 1：

输入: [1,3,5]
输出: 1
示例 2：

输入: [2,2,2,0,1]
输出: 0
说明：

这道题是 寻找旋转排序数组中的最小值 的延伸题目。
允许重复会影响算法的时间复杂度吗？会如何影响，为什么？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
fun findMin(numbers: IntArray): Int {
  if(numbers.size == 1) return numbers[0]
  if(numbers.size == 2) return kotlin.math.min(numbers[0], numbers[1])
  if (numbers[0] < numbers.last()) return numbers[0]
  if (numbers.last() < numbers[numbers.size - 2]) return numbers.last()
  var left = 0
  var right = numbers.size
  while (left < right) {
    val mid = (left + right) shr 1
    if (numbers[mid] > numbers[mid + 1]) return numbers[mid + 1]
    if (numbers[mid] < numbers[mid - 1]) return numbers[mid]
    if (numbers[mid] < numbers.last() || numbers[mid] < numbers[0]) right = mid
    else if(numbers[mid] > numbers.last() || numbers[mid] > numbers[0]) left = mid
    else return kotlin.math.min(findMin(numbers.sliceArray(0..mid)),findMin(numbers.sliceArray(mid until numbers.size)))
  }
  return numbers[left]
}

fun minArray(numbers: IntArray): Int {
  var low = 0
  var high = numbers.size - 1
  while (low < high) {
    val pivot = low + (high - low) / 2
    when {
        numbers[pivot] < numbers[high] -> high = pivot
        numbers[pivot] > numbers[high] -> low = pivot + 1
        else ->  high -= 1
    }
  }
  return numbers[low]
}
