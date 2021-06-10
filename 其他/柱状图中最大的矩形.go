package 其他

/**
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。





以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。





图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。



示例:

输入: [2,1,5,6,2,3]
输出: 10

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//利用单调栈 2次循环 找出每个位置的最左最近 和 最右最近的 小于他的值
//不容易 终于过了  看了下思路 手写的 单调栈竟然超时了。。。惨不忍睹的变量使用与多个复杂的东西
func largestRectangleArea(heights []int) int {
	var stack []int
	leftIndex := make([]int, len(heights))
	rightIndex := make([]int, len(heights))
	//storage[i] 代表i左边的最近的小于i的下标
	//左边
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			leftIndex[i] = -1
		} else {
			leftIndex[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	//右边
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) != 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			rightIndex[i] = len(heights)
		} else {
			rightIndex[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		left := leftIndex[i]
		right := rightIndex[i]
		maxArea = max(maxArea, heights[i]*(right-left-1))
	}
	return maxArea
}

//暴力  按宽度来 实时记录当前高度的最低值 面积就是宽度*高度的最低值  果断超时
//func largestRectangleArea1(heights []int) int {
//	maxArea := heights[0]
//	for i := 0; i < len(heights); i++ {
//		minHeight := heights[i]
//		for j := i; j < len(heights); j++ {
//			minHeight = min(minHeight, heights[j])
//			maxArea = max(maxArea, (j-i+1)*minHeight)
//		}
//	}
//	return maxArea
//}
//
//func min(x, y int) int {
//	if x > y {
//		return y
//	}
//	return x
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//遍历高度 一样超时
//func largestRectangleArea(heights []int) int {
//	maxArea := heights[0]
//	for i := 0; i < len(heights); i++ {
//		curHeight := heights[i] //固定高度 找出以当前高度为高 能找到的最大底部
//		//左边从i往左开始找 找到比curHeight小的停止
//		j := i
//		for j >= 0 && heights[j] >= curHeight {
//			j--
//		}
//		left := j+1
//		j = i
//		for j < len(heights) && heights[j] >= curHeight {
//			j++
//		}
//		right := j-1
//		maxArea = max(curHeight * (right - left + 1),maxArea)
//	}
//	return maxArea
//}
