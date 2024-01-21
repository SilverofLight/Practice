package main

import "fmt"
/*
题目：Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]*/
//重点是确定在那一行？
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func levelOrder(root *TreeNode) [][]int {
    //定义一个空的切片用来存储
	ret := [][]int{}
	//如果输入为空，返回空
	if root == nil {
		return ret
	}
	//q先记录第一层的数据，之后可以记录每一层有几个数据
	q := []*TreeNode{root}//len(q) = 1
	for i := 0; len(q) > 0; i++{//遍历每一行
		//在ret后面加一行
		ret = append(ret, []int{})
		//p用来暂存下一行的数量，之后交给q,每一次循环重定义为0
		p := []*TreeNode{}
		//遍历每一个数据
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	//现在的ret是从root到leaf的，需要翻转
	//可以先数出ret的行数，反向输出
	n := 0
	for i,_ := range ret {
		n = i
	}
	ret1 := [][]int{}
	for i := n; i >= 0; i--{
		ret1 = append(ret1, ret[i])
	}
	return ret1
}


func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(levelOrder(root))
}