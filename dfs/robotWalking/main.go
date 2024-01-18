package main
import "fmt"

/*一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？*/

func robotWalking(m, n int)int {
	/*实际上就是一个杨辉三角
	例如：m=7,n=3
	1  1  1  1  1  1  1 
	1  2  3  4  5  6  7
	1  3  6  10 15 21 28*/

	//定义一个数组，长和宽分别是m,n
	// ways := [][]int{}
	// dp := []int{}
	// for i := 0; i < m; i++ {
	// 	dp = append(dp, 1)
	// }
	// for i := 0; i < n; i++ {
	// 	ways = append(ways, dp)
	// }
	// fmt.Println(ways)
	//在Go语言中，当你使用append函数添加切片时，
	//如果切片底层数组容量不足，它可能会创建一个新的底层数组，
	//并将数据复制到新的数组中。这可能导致在你的代码中，所有行的切片其实都指向了同一个底层数组，而不是各自独立的数组。

	//应该使用下面的方法：
	ways := make([][]int,n)
	for i := 0; i < n; i++ {
		ways[i] = make([]int,m)
		for j := 0; j < m; j++ {
			if ways[i][j] == 0 {
				ways[i][j] = 1
			}
		}
	}
	//后面的每一个数字等于前面的数字加上面的数字
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			ways[i][j] = ways[i-1][j] + ways[i][j-1]
		}
	}
	fmt.Println(ways)
	return ways[n-1][m-1]
}
func main(){
	m := 7
	n := 4
	fmt.Println(robotWalking(m,n))
}