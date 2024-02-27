package main
import "fmt"
//方法二，使用动态规划
func minimumTotal(triangle [][]int) int {
    //如果为空，返回0
	if len(triangle) == 0 || len(triangle[0])==0 {
		return 0
	}
	//初始化dp为最后一行
	//l为最后一行长度
	l := len(triangle[len(triangle)-1])
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = triangle[len(triangle)-1][i]
	}
	//从后往前找出最小路径
	for i := len(triangle)-2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if dp[j] < dp[j+1] {
				dp[j] += triangle[i][j]
			}else {
				dp[j] = dp[j+1] + triangle[i][j]
			}
		}
	}
	return dp[0]
}
func main(){
	fmt.Println(minimumTotal([][]int{
	  []int{2},
		  []int{3, 4},
		  []int{6, 5, 7},
		  []int{4, 1, 8, 3},
	}))
  }