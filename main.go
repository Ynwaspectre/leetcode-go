package main

import (
	"fmt"
)

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

func reverseBits(num uint32) uint32 {
	var n uint32
	for i := 0; i < 32 && num > 0; i++ {
		n |= num % 2 << (31 - i) //二进制加 类似于| 0|0= 0,0|1 =1
		num >>= 1
	}
	return n
}

func main() {
	fmt.Println(reverseBits(43261596))
}
