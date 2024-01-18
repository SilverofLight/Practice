package main
import "fmt"
/*
在给定的网格中，每个单元格可以有以下三个值之一：

    值 0 代表空单元格；
    值 1 代表新鲜橘子；
    值 2 代表腐烂的橘子。

每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。

返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1
输入：[[2,1,1],[1,1,0],[0,1,1]]
输出：4
*/
func Time(res [3][3]int)int {
    t := 0
    for times := 0; times < 10; times++ {
        //数出腐烂的数量Rot和没腐烂的数量Rig
        Rot,Rig := 0,0
        a := [9]int{}
        b := [9]int{}
        count := 0 //a，b，count用来记录怀橘子的位置
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                if res[i][j] == 2 {
                    Rot++
                    a[count] = i
                    b[count] = j
                    count++
                }
                if res[i][j] == 1 {
                    Rig++
                }
            }
        }
        //如果Rig是0，返回时间t
        if Rig == 0 {
            return t
        } else {//把附近的橙子变坏,时间＋1
            for k := 0; k < Rot; k++ {
                if a[k]-1 >= 0 && res[a[k]-1][b[k]] == 1 {
                    res[a[k]-1][b[k]]=2
                }
                if b[k]-1 >= 0 && res[a[k]][b[k]-1] == 1 {
                    res[a[k]][b[k]-1]=2
                }
                if a[k] + 1 < 3 && res[a[k] + 1][b[k]] == 1 {
                    res[a[k]+1][b[k]]=2
                }
                if b[k] + 1 < 3 && res[a[k]][b[k] + 1] == 1 {
                    res[a[k]][b[k]+1]=2
                }
            }
            t++
        }
    }
    
    
    //多次都不行的话就返回-1
    return -1
    
}

func main(){
    Orianges := [3][3]int {{2,1,1},{1,1,0},{0,1,1}}
    fmt.Println(Time(Orianges))
}