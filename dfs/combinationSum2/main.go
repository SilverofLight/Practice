package main

import (
	"fmt"
	"sort"
)

/*给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。

示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]*/

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			if target-candidates[i] < 0 {
				break
			}
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))
}
