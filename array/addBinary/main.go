package main

import "fmt"

/*给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

示例 1：

输入:a = "11", b = "1"
输出："100"

示例 2：

输入：a = "1010", b = "1011"
输出："10101"



提示：

    1 <= a.length, b.length <= 104
    a 和 b 仅由字符 '0' 或 '1' 组成
    字符串如果不是 "0" ，就不含前导零
*/

func addBinary(a string, b string) string {
	//先将较短的一个前面补0，再相加
	if len(a) > len(b) {
		var c string
		for i := 0; i < len(a)-len(b); i++ {
			c += "0"
		}
		b = c + b
		//fmt.Println(b)
	} else if len(a) < len(b) {
		var c string
		for i := 0; i < len(b)-len(a); i++ {
			c += "0"
		}
		a = c + a
	}
	add := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		add[i] = int(a[i]) + int(b[i]) - 48 - 48
	}
	for i := len(a) - 1; i > 0; i-- {
		if add[i] > 1 {
			add[i] -= 2
			add[i-1] += 1
		}
	}
	if add[0] > 1 {
		c := []int{1}
		add[0] -= 2
		add = append(c, add...)

	}
	ret := ""
	for i := 0; i < len(add); i++ {
		ret += string(add[i] + 48)
	}
	return ret
}

func main() {
	fmt.Println(addBinary("111", "1"))
}
