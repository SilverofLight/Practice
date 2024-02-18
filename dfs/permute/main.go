package main

import "fmt"

/*给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：

输入：nums = [1]
输出：[[1]]


    1 <= nums.length <= 6
    -10 <= nums[i] <= 10
    nums 中的所有整数 互不相同

*/
func permute(nums []int) (ans [][]int) {
	n := len(nums)
	//path用来记录变化后的数组，ifpath用来记录某一个数是否已经用了
	path := make([]int, n)
	ifpath := make([]bool, n)
	//定义递归
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		for j, on := range ifpath {
			if !on {
				path[i] = nums[j]
				ifpath[j] = true
				dfs(i + 1)
				ifpath[j] = false
			}
		}
	}
	dfs(0)
	return
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
