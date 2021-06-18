package 二叉树

/**
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


限制：

0 <= 节点个数 <= 5000



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//前序遍历第一个元素是根节点  因为没有重复元素  所以在中序遍历里找到第一个和preorder[0]相同的元素 可以把 中序遍历和前序遍历进行划分
//然后递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil} //根节点
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//前序遍历  第1~第i个元素是 左子树 i
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
