package main

import "fmt"

/*给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：

    m == matrix.length
    n == matrix[i].length
    1 <= m, n <= 10
    -100 <= matrix[i][j] <= 100
*/

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	cil := make([]int, m*n)
	count := 0

re:
	for i := 0; ; i++ {
		for j := i; j < n-i; j++ {
			cil[count] = matrix[i][j]
			count++
			if count == m*n {
				break re
			}
		}
		for j := i + 1; j < m-i; j++ {
			cil[count] = matrix[j][n-i-1]
			count++
			if count == m*n {
				break re
			}
		}
		for j := n - i - 2; j >= i; j-- {
			cil[count] = matrix[m-1-i][j]
			count++
			if count == m*n {
				break re
			}
		}
		for j := m - i - 2; j > i; j-- {
			cil[count] = matrix[j][i]
			count++
			if count == m*n {
				break re
			}
		}
	}
	return cil
}

func main() {
	matrix := [][]int{{1, 1, 2}, {1, 1, 2}}
	fmt.Println(spiralOrder(matrix))
}
