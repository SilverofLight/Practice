package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],

	[1,2,1],
	[2,1,1]]

示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：

	1 <= nums.length <= 8
	-10 <= nums[i] <= 10
*/
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	used := make([]bool, n)
	path := make([]int, 0, n)

	var dfs func(int)
	dfs = func(idx int) {
		if idx == n {
			temp := make([]int, n)
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		for i := 0; i < n; i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(idx + 1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0)
	return
}
func main() {
	nums := []int{1, 1, 2, 1}
	fmt.Println(permuteUnique(nums))
}
