package 数组

/**

在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，
输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。



示例:

现有矩阵 matrix 如下：


给定 target = 5，返回 true。

给定 target = 20，返回 false。



限制：

0 <= n <= 1000

0 <= m <= 1000
*/

//有序数组 首先想到了二分 但是想了一会不知道咋做 突然又想到 先确定行 在确定列 然后行和列暴力。。。^_^ made和暴力没啥区别啊 看了题解原来是从右上角看
//突然想到之前遇到过 害 忘了

//看了题解shouxie
//右上角 左边的数都比他小 下面的都比他大
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
// 注意长度为0的时候 直接返回false made 这个算下来的耗时和内存竟然和上面一样
//]
func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)    //一共多少行
	col := len(matrix[0]) //一共多少列
	if row == 0 || col == 0 {
		return false
	}
	for i, j := 0, col-1; i < row && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}

//我的暴力
func findNumberIn2DArray__(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])

	targetRow := 0
	targetCol := 0

	//确定行
	for i := 0; i < row; i++ {
		if target <= matrix[i][row-1] && target >= matrix[i][0] {
			targetRow = i
		}
		return false
	}
	//确定列
	for i := 0; i < col; i++ {
		if target <= matrix[row-1][i] && target >= matrix[0][i] {
			targetCol = i
		}
		return false
	}
	return matrix[targetRow][targetCol] == target
}
