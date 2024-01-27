package main

import "fmt"
//力扣官方解法
func isValidSudoku(board [][]byte) bool {
    var rows, columns [9][9]int
    var subboxes [3][3][9]int//三维数组
    for i, row := range board {
        for j, c := range row {
            if c == '.' {
                continue
            }
            index := c - '1'
            rows[i][index]++
            columns[j][index]++
            subboxes[i/3][j/3][index]++
            if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
                return false
            }
        }
    }
    return true
}

/*作者：力扣官方题解
链接：https://leetcode.cn/problems/valid-sudoku/solutions/1001859/you-xiao-de-shu-du-by-leetcode-solution-50m6/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。*/