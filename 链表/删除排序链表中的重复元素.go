package 链表

/**
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。



示例 1：


输入：head = [1,1,2]
输出：[1,2]
示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//
// 虽然是简单 但是对链表并不是很熟悉，不过还是自己做出来了，本来想用个map来保存遇到的 然后遍历的时候map
//里有就删除 后来看是有序的应该可以连续的
func deleteDuplicates(head *ListNode) *ListNode {
	head1 := head //先找个替代的去遍历
	for head1 != nil && head1.Next != nil {
		if head1.Next.Val == head1.Val {
			head1.Next = head1.Next.Next
		} else { //刚开始这边没加else 没通过 只会对1个重复的进行删除 后面加了else后 有重复的不会向前走，就会一直删除重复的 比较关键
			head1 = head1.Next
		}
	}
	return head
}
