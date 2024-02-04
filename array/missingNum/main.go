package main

import "fmt"

/*给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：

输入：nums = [1,2,0]
输出：3

示例 2：

输入：nums = [3,4,-1,1]
输出：2

示例 3：

输入：nums = [7,8,9,11,12]
输出：1



提示：

    1 <= nums.length <= 5 * 105
    -231 <= nums[i] <= 231 - 1
*/

//本题可以先排序再查找即可，这个方法就不写了，写一个力扣官方的用哈希表的方法

//由于长度为N的数列，缺失的数字一定在【1，N+1】之间，
//1.先把负数改成范围之外的数字
//2.把每一位数字-1的位置改为负数
//3.第一个不是负数的位置+1就是所求

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= len(nums) {
			nums[nums[i]+1] = -1
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
