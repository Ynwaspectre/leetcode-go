package 二叉树

/**
给定一个二叉树的根节点 root ，返回它的 中序 遍历。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}
