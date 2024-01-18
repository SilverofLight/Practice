package main
import (
	"fmt"
)
/*
题目：对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。

示例 1：

输入：str1 = "ABCABC", str2 = "ABC"
输出："ABC"
示例 2：

输入：str1 = "ABABAB", str2 = "ABAB"
输出："AB"
示例 3：

输入：str1 = "LEET", str2 = "CODE"
输出：""
*/
func gcd(a int, b int)int {
	//求两个字符串长度的最大公因数
	var c int
	for i := 1; i <= a && i <= b; i++{
		if a % i == 0 && b % i == 0 {
			c = i
		}
	}
	return c
}

func gcdString(s1 string, s2 string)string{
	//先求最大公因数，如果这个长度的字符串是所求就返回，否则返回“”
	length := gcd(len(s1), len(s2))//最大公因数
	for i := 0; i < len(s1); i++{
		if s1[i] != s1[i-i/length*length]{
			return ""
		}
	}
	for i := 0; i < len(s2); i++{
		if s2[i] != s2[i-i/length*length]{
			return ""
		}
	}
	return s1[ : length]
}

func main(){
	fmt.Println(gcdString("ABABA","ABAB"))
}