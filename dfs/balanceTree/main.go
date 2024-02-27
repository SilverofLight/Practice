package main

/*判断一棵树是否是平衡树,如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树
Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.*/
//本题仍然使用递归的方法解题

type TreeNode struct { //定义一个二叉树的结构体
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//需要判断每一个节点是否是平衡
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//递归，从上往下判断是否是平衡点
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//递归，数出某个节点的高度
	max(height(root.Left), height(root.Right)) + 1
}

//max函数返回较大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//abs函数用来计算绝对值
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {

}
