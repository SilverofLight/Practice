package main

import "fmt"
/*
题目：给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:
可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
*/
func lengthOfLIS(nums []int)int {
	length := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		count++
		if nums[i + 1] < nums[i] && count > length {
			length = count+1
			count = 0
		}
	}
	return length
}
func main(){
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}