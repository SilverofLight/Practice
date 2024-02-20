package main

import (
	"fmt"
	"strings"
)

/*按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。

示例 2：

输入：n = 1
输出：[["Q"]]

提示：

    1 <= n <= 9
*/
// func solveNQueens(n int) [][]string {
// 	queen := make([]string, n)
// 	var solve [][]string
// 	var on [][]int
// 	for i := 0; i < n; i++ {
// 		on = append(on, []int(nil))
// 		for j := 0; j < n; j++ {
// 			on[i] = append(on[i], 0)
// 		}
// 	}
// 	num := 0

// 	var dfs func(int, int)
// 	dfs = func(i, n int) {
// 		if i == n {
// 			settle(on, queen, n)
// 			solve = append(solve, append([]string(nil), queen...))
// 			return
// 		}
// 		for j := 0; j < n; j++ {
// 			if on[i][j] == 0 {
// 				occupy(n, on, i, j)
// 				num++
// 				dfs(i+1, n)
// 				undo(n, on, i, j)
// 				num--
// 			}
// 		}
// 	}
// 	dfs(0, n)
// 	return solve
// }

// //用来判断那个位置不能放Q
// func occupy(n int, on [][]int, i int, j int) {
// 	for l := 0; l < n; l++ {
// 		on[i][l] = 1
// 		on[l][j] = 1
// 	}
// 	for l := 0; l+i < n && l+j < n; l++ {
// 		on[l+i][l+j] = 1
// 	}
// 	for l := 0; i-l >= 0 && j-l >= 0; l++ {
// 		on[i-l][j-l] = 1
// 	}
// 	on[i][j] = 2
// }

// //撤销occupy
// func undo(n int, on [][]int, i int, j int) {
// 	for l := 0; l < n; l++ {
// 		on[i][l] = 0
// 		on[l][j] = 0
// 	}
// 	for l := 0; l+i < n && l+j < n; l++ {
// 		on[l+i][l+j] = 0
// 	}
// 	for l := 0; i-l >= 0 && j-l >= 0; l++ {
// 		on[i-l][j-l] = 0
// 	}
// }

// //把on数组转化成queen
// func settle(on [][]int, queen []string, n int) {
// 	m := map[int]string{
// 		2: "Q",
// 		1: ".",
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			queen[i] += m[on[i][j]]
// 		}
// 	}
// }

func solveNQueens(n int) (ans [][]string) {

	col := make([]int, n)

	onPath := make([]bool, n)

	diag1 := make([]bool, n*2-1)

	diag2 := make([]bool, n*2-1)

	var dfs func(int)

	dfs = func(r int) {

		if r == n {

			board := make([]string, n)

			for i, c := range col {

				board[i] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-1-c)

			}

			ans = append(ans, board)

			return

		}

		for c, on := range onPath {

			rc := r - c + n - 1

			if !on && !diag1[r+c] && !diag2[rc] {

				col[r] = c

				onPath[c], diag1[r+c], diag2[rc] = true, true, true

				dfs(r + 1)

				onPath[c], diag1[r+c], diag2[rc] = false, false, false // 恢复现场

			}

		}

	}

	dfs(0)

	return

}

// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/n-queens/solutions/2079586/hui-su-tao-lu-miao-sha-nhuang-hou-shi-pi-mljv/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func main() {
	fmt.Println(solveNQueens(3))
}

//打开力扣->看题->翻答案->提交
