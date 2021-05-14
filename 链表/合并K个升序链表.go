package 链表

import (
	"sort"
)

/**
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//^_^ 比较暴力
//看了好像是用什么堆 无奈没咋用过 先这样吧
//或者额可以用之前的合并2个的 每2个合并
func mergeKLists(lists []*ListNode) *ListNode {
	var nums []int
	for _, v := range lists {
		for v != nil {
			nums = append(nums, v.Val)
			v = v.Next
		}
	}
	sort.Ints(nums)
	if len(nums) == 0 {
		return nil
	}
	return createNode(nums)
}

//func createNode(nums []int) *ListNode {
//	head := &ListNode{nums[0], nil}
//	head1 := head
//	for i := 1; i < len(nums); i++ {
//		head1.Next = &ListNode{nums[i], nil}
//		head1 = head1.Next
//	}
//	return head
//}
