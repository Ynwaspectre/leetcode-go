package 数组

import "math"

/**
回旋镖定义为一组三个点，这些点各不相同且不在一条直线上。

给出平面上三个点组成的列表，判断这些点是否可以构成回旋镖。



示例 1：

输入：[[1,1],[2,3],[3,2]]
输出：true
示例 2：

输入：[[1,1],[2,2],[3,3]]
输出：false


提示：

points.length == 3
points[i].length == 2
0 <= points[i][j] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-boomerang
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//咋一看判断是否能变成三角形
func isBoomerang(points [][]int) bool {
	//先判断前2个吧
	a := points[0]
	b := points[1]
	c := points[2]

	//1.不能在一条线

	//2.不能是相同的点

	return true
}
