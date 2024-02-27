package main

import "fmt"

/*给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

输入：head = [0,1,2], k = 4
输出：[2,0,1]


    链表中节点的数目在范围 [0, 500] 内
    -100 <= Node.val <= 100
    0 <= k <= 2 * 109

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	//需要将链表合并成环之后旋转
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	next := head
	for next.Next != nil {
		n++
		next = next.Next
	}
	//n求出了链表的数量
	move := n - k%n
	if move == n {
		return head
	}
	next.Next = head
	for move > 0 {
		next = next.Next
		move--
	}
	ret := next.Next
	//将环打开
	next.Next = nil
	return ret
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.Next = &ListNode{Val: arr[i]}
		current = current.Next
	}
	return head
}

func main() {
	// Create and initialize the linked list
	arr := []int{1, 2, 3, 4, 5}
	head := createList(arr)

	// Rotate the linked list
	k := 2
	rotated := rotateRight(head, k)

	// Print the rotated linked list
	printList(rotated) // Output: 4 5 1 2 3
}
