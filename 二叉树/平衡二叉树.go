package 二叉树

import "math"

/**
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。



限制：

0 <= 树的结点个数 <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//这个解法。。确实有重复的 做的时候就觉得了 我的想法是自底向上 但是初始是头结点
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(float64(getHeight(root.Left)-getHeight(root.Right))) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deepL := getHeight(root.Left)
	deepR := getHeight(root.Right)
	if deepL > deepR {
		return deepL + 1
	}
	return deepR + 1
}
