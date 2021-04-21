package 链表

/**
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//思路比较麻烦 就是把链表中元素都拿出来 放到数组里 然后翻转数组重新生成列表 嘿嘿
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var nums []int
	head1 := head
	for head1 != nil {
		nums = append(nums, head1.Val)
		head1 = head1.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		temp := nums[i]
		nums[i] = nums[len(nums)-i-1]
		nums[len(nums)-i-1] = temp
	}
	return createNode(nums)
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
