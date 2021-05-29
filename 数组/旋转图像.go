package 数组

/**
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//[0][0]=>[0][3]
// [0][1]=>[1][3]
//

//拷贝了一份 然后把那一份按列循环 赋值到目前的行上
func rotate(matrix [][]int) {
	copyMatrix := make([][]int, len(matrix))
	copy(copyMatrix, matrix)
	for i := 0; i < len(matrix[0]); i++ {
		temp := make([]int, len(matrix[0]))
		for j := len(matrix) - 1; j >= 0; j-- {
			temp[len(matrix)-j-1] = copyMatrix[j][i]
		}
		matrix[i] = temp
	}
}
