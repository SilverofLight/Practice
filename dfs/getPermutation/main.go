package main

import "fmt"

/*给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

    "123"
    "132"
    "213"
    "231"
    "312"
    "321"

给定 n 和 k，返回第 k 个排列。

示例 1：

输入：n = 3, k = 3
输出："213"

示例 2：

输入：n = 4, k = 9        1234, 1243, 1324, 1342, 1423, 1432,||| 2134, 2143, 2314, 2341,
输出："2314"			第一步：(k - 1) / (n - 1)!,   8 / 6 = 1, 所以要找的数在2开头的数字里面， 把2从数组里去掉，str首位加上2
						第二步：(k - 1) % (n - 1)!,  8 % 6 = 2, 所以要找的数在的二组的第三个
示例 3：				第三步：2 / (n - 2)!, 2 / 2 = 1, 在第二组里，1,3,4中把3去掉，str后面加上3
						第四步：2 % (n - 2)!, 2 % 2 = 0, 在这一组的第一个，直接把剩下的加到str后面

输入：n = 3, k = 1
输出："123"

提示：

    1 <= n <= 9
    1 <= k <= n!
*/
//思路是直接找出需要的组合是哪个

func getPermutation(n int, k int) string {
	//初始化
	nums := make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = byte('1' + i)
	}
	//fmt.Println(nums)
	var str string

	//定义递归
	var dfs func(int, int)
	dfs = func(ni int, ki int) {
		sa := ki / factorial(ni)
		str += string(nums[sa])
		var s string
		for j := 0; j < len(nums); j++ {
			if j == sa {
				continue
			}
			s += string(nums[j])
		}
		for j := 0; j < len(nums)-1; j++ {
			nums[j] = s[j]
		}
		nums = nums[:len(nums)-1]
		if ki%factorial(ni) == 0 {
			str += btoS(nums)
			return
		}
		//fmt.Println(nums)
		dfs(ni-1, ki%factorial(ni))
	}
	dfs(n-1, k-1)
	return str
}

//求阶乘
func factorial(n int) int {
	m := 1
	for i := 1; i <= n; i++ {
		m *= i
	}
	return m
}

//给一个byte组，去掉其中的一位,返回byte组
// func del(nums []byte, i int) []byte {
// 	var s string
// 	for j := 0; j < len(nums); j++ {
// 		if j == i {
// 			continue
// 		}
// 		s += string(nums[j])
// 	}
// 	for j := 0; j < len(nums)-1; j++ {
// 		nums[j] = s[j]
// 	}
// 	nums = nums[:len(nums)-1]
// 	return nums
// }

//给一个byte组，转换成string
func btoS(b []byte) string {
	var s string
	for i := 0; i < len(b); i++ {
		s += string(b[i])
	}
	return s
}

func main() {
	fmt.Println(getPermutation(3, 3))
}
