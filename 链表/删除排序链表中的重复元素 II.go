package 链表

/**
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。

返回同样按升序排列的结果链表。



示例 1：


输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：


输入：head = [1,1,1,2,3]
输出：[2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//跟1的区别在于不保留重复元素 而且也变成中等难度了
//还是自己想出来了 一次ac，上一个删除是保留 所以可以从遇到的第一个重复的开始，这次是都要删除，
//先弄了个map记录了下当前数值的数量情况，因为没办法解决刚开始就遇到重复的数的情况，所以就在原先
//节点前面加了个0节点，数字为所谓next为head 从这个节点开始遍历，最后指向这个节点的next就是要返回的
//这个数字有几个就next几次 一个循环

//对比官方解法 果然是要在head前增加一个节点 ，官方叫哑节点，但是官方没有先遍历一遍存下数据 就直接一次遍历
func deleteDuplicatesII(head *ListNode) *ListNode {
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
