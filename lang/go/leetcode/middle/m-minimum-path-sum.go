package leetcode

/*
*
64. 最小路径和

给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
[1,3,1],
[1,5,1],
[4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var dp = make([]int, len(grid[0]))

	for x := len(grid) - 1; x >= 0; x-- {
		for y := len(grid[0]) - 1; y >= 0; y-- {
			if x == len(grid)-1 {
				//原本应该首先单独判断的，放进来了，代码会简洁，可读性会变差
				if y == len(grid[0])-1 {
					dp[y] = grid[len(grid)-1][len(grid[0])-1]
				} else {
					dp[y] = dp[y+1] + grid[x][y]
				}
				continue
			}
			//原本应该分开写的
			if y == len(grid[0])-1 {
				dp[y] = dp[y] + grid[x][y]
			} else {
				dp[y] = min(dp[y], dp[y+1]) + grid[x][y]
			}
		}
	}
	return dp[0]
}

type Point struct {
	x int
	y int
}

// dfs超时
func minPathSumV2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var m = make(map[Point]int)
	return minPathSumV2Dfs(Point{0, 0}, m, grid, 0)
}

func minPathSumV2Dfs(p Point, m map[Point]int, grid [][]int, sum int) int {
	if ret, ok := m[p]; ok {
		return ret
	}
	var ret int
	if p.x == len(grid)-1 {
		if p.y == len(grid[0])-1 {
			ret = sum + grid[p.x][p.y]
		} else {
			ret = minPathSumV2Dfs(Point{p.x, p.y + 1}, m, grid, sum) + grid[p.x][p.y]
		}
		m[p] = ret
		return ret
	}
	if p.y == len(grid[0])-1 {
		ret = minPathSumV2Dfs(Point{p.x + 1, p.y}, m, grid, sum) + grid[p.x][p.y]
	} else {
		ret = min(minPathSumV2Dfs(Point{p.x + 1, p.y}, m, grid, sum), minPathSumV2Dfs(Point{p.x, p.y + 1}, m, grid, sum)) + grid[p.x][p.y]
	}
	m[p] = ret
	return ret
}
