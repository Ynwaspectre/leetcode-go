package 其他

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：



输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9


提示：

n == height.length
0 <= n <= 3 * 104
0 <= height[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//看的题解动态规划  不咋需要看代码 只要看思路就可以写出来了
//最关键的一点是 第54行 第i行的最大盛水是 左边的最大高度和右边的最大高度 2个值取最小 然后减去自己的高度
func trap(height []int) int {
	length := len(height)

	if length == 0 {
		return 0
	}

	leftMax := make([]int, length)
	rightMax := make([]int, length)

	leftMax[0] = height[0]
	for i := 1; i < length; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	rightMax[length-1] = height[length-1]
	for i := length - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	total := 0
	for i := 0; i < length; i++ {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}
	return total
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
