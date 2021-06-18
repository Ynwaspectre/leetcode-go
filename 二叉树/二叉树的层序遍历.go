package 二叉树

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
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
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var stack []*TreeNode
	if root == nil {
		return result
	}
	stack = append(stack, root)
	for len(stack) > 0 {
		//全部pop出来处理
		var floor []int
		var insert []*TreeNode
		for len(stack) > 0 {
			floor = append(floor, stack[len(stack)-1].Val)
			insert = append(insert, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		result = append(result, floor)
		for len(insert) > 0 {
			temp := insert[len(insert)-1]
			if temp.Right != nil {
				stack = append(stack, temp.Right)
			}
			if temp.Left != nil {
				stack = append(stack, temp.Left)
			}
			insert = insert[:len(insert)-1]
		}
	}
	return result
}
