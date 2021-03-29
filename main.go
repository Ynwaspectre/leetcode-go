package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(nums []int) *ListNode {
	head := &ListNode{nums[0], nil}
	head1 := head
	for i := 1; i < len(nums); i++ {
		head1.Next = &ListNode{nums[i], nil}
		head1 = head1.Next
	}
	return head
}

func printNode(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	head1 := head
	head2 := &ListNode{0, head}
	head3 := head2
	m := make(map[int]int)
	for head1 != nil {
		m[head1.Val]++
		head1 = head1.Next
	}
	for head2.Next != nil {
		if v, _ := m[head2.Next.Val]; v > 1 {
			head2.Next = head2.Next.Next
			v--
		} else {
			head2 = head2.Next
		}
	}
	return head3.Next
}

func main() {
	nums := createNode([]int{1, 1})
	node := deleteDuplicates(nums)
	printNode(node)

}
