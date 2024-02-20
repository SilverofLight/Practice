package main

import "fmt"

/*实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000

示例 2：

输入：x = 2.10000, n = 3
输出：9.26100

示例 3：

输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25

提示：

    -100.0 < x < 100.0
    -2^31 <= n <= 2^31-1
    n 是一个整数
    要么 x 不为零，要么 n > 0 。
    -10^4 <= x^n <= 10^4
*/

//笨方法，一次一次乘，超时了
// func myPow(x float64, n int) float64 {
// 	pow := x
// 	if n > 0 {
// 		for i := 1; i < n; i++ {
// 			pow = pow * x
// 		}
// 	} else if n == 0 {
// 		pow = 1
// 	} else {
// 		for i := 1; i > n; i-- {
// 			pow = pow / x
// 		}
// 	}
// 	return pow
// }

func main() {
	fmt.Println(myPow(100, -1))
}

func myPow(x float64, n int) float64 {
	pow := x
	//x -> x^2 -> x^4 -> x^8
	var dfs func(float64, int) float64
	dfs = func(f float64, i int) float64 {
		if i == 0 {
			return 1
		}
		y := dfs(f, i/2)
		if i%2 == 0 {
			return y * y
		}
		return y * y * f
	}
	if n > 0 {
		pow = dfs(x, n)
	} else {
		pow = 1.0 / dfs(x, -n)
	}
	return pow
}
