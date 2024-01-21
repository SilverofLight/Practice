package main

import "fmt"

/*题目：
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//使用同步移动的方法，即定义p和q，p左移q就右移，判断二者是否相等
func Mirror(root *TreeNode)bool{
	if root == nil {
		return true
	}
	return check(root, root)
}
func check(p,q *TreeNode)bool{
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p == q && check(q.Right,p.Left) && check(q.Left, p.Right)
}

func main(){
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(Mirror(root))
}