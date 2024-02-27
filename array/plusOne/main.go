package main

import "fmt"

/*给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。

示例 3：

输入：digits = [0]
输出：[1]



提示：

    1 <= digits.length <= 100
    0 <= digits[i] <= 9

*/

//需要找出末尾有几个9
func plusOne(digits []int) []int {
	numOf9 := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			numOf9++
		} else {
			break
		}
	}
	if numOf9 == 0 {
		digits[len(digits)-1]++
		return digits
	}
	if numOf9 == len(digits) {
		digits = append(digits, 0)
		digits[0] = 1
		for i := 1; i < len(digits); i++ {
			digits[i] = 0
		}
		return digits
	}
	digits[len(digits)-1-numOf9]++
	long := len(digits)
	for numOf9 > 0 {
		digits[long-1] = 0
		long--
		numOf9--
	}
	return digits
}
func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}
