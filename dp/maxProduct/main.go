package main
import (
	"fmt"
	"math"
)

/*
题目：Given an integer array nums, find the contiguous subarray within an array (containing at least one number) 
which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/
//由于负负得正，需要保存最小的乘积和最大的乘积以备后期翻转。
func maxProduct(nums []int)int{
	//如果数组为空，返回0
	if len(nums) == 0 {
		return 0
	}
	//用dp记录每一个以nums[i]开头，到指定结尾的乘积
	dp := [][]int{
		[]int{nums[0],nums[0]},
	}
	//把每一个乘积加到dp后面
	for i := 1; i < len(nums); i++ {
		maxNum := int(math.Max(float64(nums[i]*dp[i-1][0]), float64(nums[i]*dp[i-1][1])))
		minNum := int(math.Min(float64(nums[i]*dp[i-1][0]), float64(nums[i]*dp[i-1][1])))
		dp = append(dp, []int{
			int(math.Max(float64(nums[i]), float64(maxNum))),
			int(math.Min(float64(nums[i]), float64(minNum))),
		})
	}
	//遍历一遍dp，找出最大的数
	res := dp[0][0]
	for i := 1; i < len(dp); i++{
		if res < dp[i][0]{
			res = dp[i][0]
		}
	}
	return res
}

func main(){
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}