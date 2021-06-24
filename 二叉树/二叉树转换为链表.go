package 二叉树

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：


输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [0]
输出：[0]


提示：

树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100


进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
感觉还是递归吧
*/

//思路 先整理左边 然后整理右边 最后把左边的接到右边去即可
//还有优化空间 比如新加个函数 带返回值的 这样遍历到结束就可以把最后一个节点的地址知道了 就不用后续再去了
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left == nil {
		return
	}
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return
	}
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	head := root.Right
	for head.Right != nil {
		head = head.Right
	}
	head.Right = temp
}
