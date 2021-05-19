package 二叉树

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//虽然是简单的 但是之前没写过二叉树 想了一会写出来了 递归 类似于多个子问题变成最后的问题
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	Val := 0
	if root1 != nil && root2 != nil {
		Val += root1.Val
		Val += root2.Val
	} else if root1 != nil {
		Val += root1.Val
	} else if root2 != nil {
		Val += root2.Val
	} else {
		return nil
	}

	newTreeNode := &TreeNode{Val, nil, nil}
	if root1 != nil && root2 != nil {
		newTreeNode.Left = mergeTrees(root1.Left, root2.Left)
		newTreeNode.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 != nil {
		newTreeNode.Left = mergeTrees(root1.Left, nil)
		newTreeNode.Right = mergeTrees(root1.Right, nil)
	} else if root2 != nil {
		newTreeNode.Left = mergeTrees(nil, root2.Left)
		newTreeNode.Right = mergeTrees(nil, root2.Right)
	}
	return newTreeNode
}
