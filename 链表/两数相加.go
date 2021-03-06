package 链表

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。



示例 1：


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]


提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//只想到一个思路 先把两个数表示出来 然后对两个数相加 然后在表示回去 但是这也太。。然后觉得可以从头开始
//按照前面开始对齐 有进位就记住 到下一个的时候加上 如果2个链表都空了要看下当前是否有进位 有的话 就next个节点 值为1
//看了下题解 和官方思路一样 ！自己想出来了 牛！但是63~70行总感觉又点累赘的感觉
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	l4 := l3
	jinwei := 0
	for l1 != nil || l2 != nil {
		value := jinwei
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		jinwei = 0
		if value >= 10 {
			jinwei++ //下一步是否需要进位
			l3.Val = value % 10
		} else {
			l3.Val = value
		}
		if l1 == nil && l2 == nil {
			if jinwei == 1 {
				l3.Next = &ListNode{1, nil}
			}
		} else {
			l3.Next = &ListNode{0, nil}
			l3 = l3.Next
		}
	}
	return l4
}
