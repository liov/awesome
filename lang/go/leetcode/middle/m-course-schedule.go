package leetcode

/**
207. 课程表

你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//一遍过，哈哈哈
//执行用时： 1336 ms , 在所有 Kotlin 提交中击败了 7.69% 的用户
//内存消耗： 54.9 MB , 在所有 Kotlin 提交中击败了 100.00% 的用户
//java 版本467ms，难道kotlin真比java慢3倍
func canFinish(numCourses int, prerequisites [][]int) bool {
	var m = make(map[int]map[int]struct{})
	for _, arr := range prerequisites {
		if _, ok := m[arr[1]]; ok {

			m[arr[1]][arr[0]] = struct{}{}
		} else {
			m2 := make(map[int]struct{})
			m2[arr[0]] = struct{}{}
			m[arr[1]] = m2
		}
	}
	var set = make(map[int]struct{})
	var mem = make(map[int]struct{})
	for k, _ := range m {
		if !canFinishDfs(k, set, mem, m) {
			return false
		}
	}
	return true
}

func canFinishDfs(num int, set map[int]struct{}, mem map[int]struct{}, m map[int]map[int]struct{}) bool {
	if _, ok := set[num]; ok {
		return false
	}
	if _, ok := mem[num]; ok {
		return true
	}
	set[num] = struct{}{}
	if _, ok := m[num]; ok {
		for n := range m[num] {
			if _, ok := mem[n]; ok {
				return true
			}
			if !canFinishDfs(n, set, mem, m) {
				return false
			}
		}
	}
	delete(set, num)
	return true
}

var canFinishValid = true

func canFinishV2(numCourses int, prerequisites [][]int) bool {
	var edges = make([][]int, 0)
	for _ = range numCourses {
		edges = append(edges, make([]int, 0))
	}
	var visited = make([]int, numCourses)
	for _, arr := range prerequisites {
		edges[arr[1]] = append(edges[arr[1]], arr[0])
	}
	var i = 0
	for i < numCourses && canFinishValid {
		if visited[i] == 0 {
			canFinishV2Dfs(i, visited, edges)
		}
		i++
	}
	return canFinishValid
}

func canFinishV2Dfs(u int, visited []int, edges [][]int) {
	visited[u] = 1
	for v := range edges[u] {
		if visited[v] == 0 {
			canFinishV2Dfs(v, visited, edges)
			if !canFinishValid {
				return
			}
		} else if visited[v] == 1 {
			canFinishValid = false
			return
		}
	}
	visited[u] = 2
}
