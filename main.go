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
	for head1 != nil && head1.Next != nil {
		if head1.Next.Val == head1.Val {
			head1.Next = head1.Next.Next
		} else {
			head1 = head1.Next
		}
	}
	return head
}

func main() {
	nums := []int{1, 2, 2, 2, 2, 2, 2, 3, 4, 5, 5}
	node := createNode(nums)
	printNode(node)
	fmt.Println("---")

	printNode(deleteDuplicates(node))
}
