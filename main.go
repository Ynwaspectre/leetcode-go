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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	l4 := l3
	jinwei := 0
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + jinwei
		jinwei = 0
		if value >= 10 {
			jinwei++ //下一步是否需要进位
			l3.Val = value % 10
		} else {
			l3.Val = value
		}
		l1 = l1.Next
		l2 = l2.Next
		l3.Next = &ListNode{0, nil}
		l3 = l3.Next
	}
	if l1 == nil && l2 != nil {
		for l2 != nil {
			value := l2.Val + jinwei
			if value > 10 {
				jinwei++ //下一步是否需要进位
				l3.Val = value % 10
			} else {
				l3.Val = value
			}
			l2 = l2.Next
			l3 = l3.Next
			jinwei = 0
		}
	} else if l2 == nil && l1 != nil {
		for l1 != nil {
			value := l1.Val + jinwei
			if value > 10 {
				jinwei++ //下一步是否需要进位
				l3.Val = value % 10
			} else {
				l3.Val = value
			}
			l1 = l1.Next
			l3 = l3.Next
			jinwei = 0
		}
	} else {
		if jinwei == 1 {
			l3 = &ListNode{1, nil}
		} else {
			l3 = nil
		}
	}
	return l4
}

func main() {
	//l1 := createNode([]int{2, 4, 3})
	//l2 := createNode([]int{5, 6, 4})
	l1 := createNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := createNode([]int{9, 9, 9, 9})
	node := addTwoNumbers(l1, l2)
	printNode(node)

}
