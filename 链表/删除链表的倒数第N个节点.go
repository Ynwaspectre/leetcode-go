package 链表

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//感觉在哪看过  设置一个步伐长度为N的2个指针
//还算比较快的解决了，在链表长度较短的时候需要做下处理
//链表问题舍得定义变量问题不大！
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//2个节点 其中一个先走n步
	resultHead := head
	head1 := head
	head2 := head
	for head2 != nil && n > 0 { //head2先走n步
		head2 = head2.Next
		n--
	}
	//一起
	for head2 != nil && head2.Next != nil {
		head2 = head2.Next
		head1 = head1.Next
	}
	if (head1 == head) && head2 == nil {
		//没动
		resultHead = resultHead.Next
	} else {
		head1.Next = head1.Next.Next
	}
	return resultHead
}
