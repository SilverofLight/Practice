package main
import "fmt"
/*
题目：给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
*/
func find(nums []int)int {
	m := make(map[int]int)

	for _,n := range nums {
		m[n]++
		if m[n] == len(nums)/2+1 {
			return n
		}
	}	
	return -1
}
func main(){
	fmt.Println(find([]int{2,2,1,1,1,2,2}))
}