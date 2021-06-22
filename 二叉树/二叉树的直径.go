package 二叉树

/**
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
这条路径可能穿过也可能不穿过根结点。



示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。



注意：两结点之间的路径长度是以它们之间边的数目表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//遍历 查看每个节点2边的深度加起来的值最大
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	//两边的加起来的最大值
	var deep func(node *TreeNode) int
	ans := 0
	deep = func(node *TreeNode) int { //deep函数计算 这个节点的最大深度 计算的过程中 计算直径 每计算一个节点 顺便计算直径
		//保留最大直径
		if root == nil {
			return 0
		}
		l := deep(node.Left)
		r := deep(node.Right)
		ans = max(l+r+1, ans)
		return max(l, r) + 1
	}
	deep(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
