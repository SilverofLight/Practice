package main
import "fmt"
/*
题目：给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]

自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。*/

//使用递归的方法返回最小的路径和

//用这个函数接收路径和
func minimumTotal(a [][]int)int {
  //如果输入的数组为空，返回0
	if len(a) == 0 || len(a[0]) == 0 {
		return 0
	}
	//定义一个变量为首行的数字
	sum := a[0][0]
	sum += min(a, 1, 0)
	return sum
}

//定义一个递归函数
func min(a [][]int, i, j int)int {
	//设置终止的条件，当i大于等于行数时终止，返回0
	if i >= len(a){
		return 0
	}
	//开始递归，用a0和a1分别接收下一行两个相邻的数字
	a0 := min(a, i+1, j)
	a1 := min(a, i+1, j+1)

	//判断和较小的路径返回
	//fmt.Printf("a0=%v, a1=%v \n",a0,a1)
	if a[i][j]+a0 < a1+a[i][j+1] {
        return a[i][j] + a0
    }
    return a[i][j+1] + a1
}

func main(){
  fmt.Println(minimumTotal([][]int{
    []int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
  }))
}