package main

//自己写一遍

func combinationSum(comparable []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(comparable) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		dfs(target, idx+1)
		if target-comparable[idx] >= 0 {
			comb = append(comb, comparable[idx])
			dfs(target-comparable[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
