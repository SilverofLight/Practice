package main

import "fmt"

/*给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：

输入：height = [4,2,0,3,2,5]
输出：9
提示：

    n == height.length
    1 <= n <= 2 * 104
    0 <= height[i] <= 105
详细：  https://leetcode.cn/problems/trapping-rain-water/description/
*/

//向右向左遍历两次，选取两次遍历中较小的

func trap(height []int) int {
	max := 0
	left := []int{}
	right := []int{}
	sum := 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
		right = append(right, max-height[i])
	}
	max = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > max {
			max = height[i]
		}
		left = append(left, max-height[i]) //left是倒序的
	}
	for i := 0; i < len(height); i++ {
		sum += min(right[i], left[len(left)-1-i])
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
