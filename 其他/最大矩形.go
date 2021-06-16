package 其他

/**
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。



示例 1：


输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = []
输出：0
示例 3：

输入：matrix = [["0"]]
输出：0
示例 4：

输入：matrix = [["1"]]
输出：1
示例 5：

输入：matrix = [["0","0"]]
输出：0


提示：

rows == matrix.length
cols == matrix[0].length
0 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//在看了题解一没看懂后 看了一个java题解后 我准备动手了！
//题解1是计算左边的  那我就来计算右边的 gg 发现并走不通 我sb了
//又去看了下 明白了 其实就是把每个点作为矩形的最右下角的点 然后把这个点能行成的最大宽度预先计算出来
//name这个点所能构成的面积就是 w*j 预先宽度*高度

/**
[["1","0","1","1","1"],["0","1","0","1","0"],["1","1","0","1","1"],["1","1","0","1","1"],["0","1","1","1","1"]]

这个案例没跑过
*/

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	leftNum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		leftNum[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if j == 0 {
					leftNum[i][j] = 1
				} else {
					leftNum[i][j] = leftNum[i][j-1] + 1
				}
			}
		}
	}
	ans := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' { //是0 不能组成
				continue
			}
			w := leftNum[i][j]
			area := w
			for k := i - 1; k >= 0; k-- {
				w = min(w, leftNum[k][j]) //made 这边原本写错了 写成重新赋值了 结果搞了半天都不对
				area = max(area, w*(i-k+1))
			}
			ans = max(ans, area)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
