package 二叉树

import "fmt"

/**
依稀记得以前有前序遍历，中序遍历，后序遍历
//前序   根-左节点-右键点
//中序   左节点-根-右节点
//后序  左节点-右节点-根节点
//层次遍历  从上到下依次
*/

//按前序遍历构建
func prev(node *TreeNode) {
	if node != nil {
		fmt.Println(node.Val)
		prev(node.Left)
		prev(node.Right)
	}
}

//中序遍历 左 根 右
func mid(node *TreeNode) {
	if node != nil {
		mid(node.Left)
		fmt.Println(node.Val)
		mid(node.Right)
	}
}

func end(node *TreeNode) {
	if node != nil {
		end(node.Left)
		end(node.Right)
		fmt.Println(node.Val)
	}
}

//放到队列
func cengxu(node *TreeNode) {
	var nodeQueue []*TreeNode
	nodeQueue = append(nodeQueue, node)
	for len(nodeQueue) > 0 {
		currNode := nodeQueue[len(nodeQueue)-1]
		fmt.Println(currNode.Val)
		nodeQueue = nodeQueue[:len(nodeQueue)-1]
		if currNode.Left != nil {
			nodeQueue = append([]*TreeNode{currNode.Left}, nodeQueue...)
		}
		if currNode.Right != nil {
			nodeQueue = append([]*TreeNode{currNode.Right}, nodeQueue...)
		}
	}
}
