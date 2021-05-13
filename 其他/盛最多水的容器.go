package 其他

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。



示例 1：



输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
示例 3：

输入：height = [4,3,2,1,4]
输出：16
示例 4：

输入：height = [1,2,1]
输出：2


提示：

n = height.length
2 <= n <= 3 * 104
0 <= height[i] <= 3 * 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//写思路最重要了 这题提交成功率挺高呀为什么我就是想不出啥思路
// 想过双指针从2边找没去尝试   看了相关提示双指针关键词 没看答案 来做做看 一次ok 但是时间很长
// 其实感觉有很多可以省略 假设找到2个最大的了 那中间的就没必要了
//看了题解 当等于的时候好像随便移动 这个在我看来比较难理解 我刚开始的理解
//终于找到有个题解说这个问题的了
//原因：要想找到比当下还要大的值的时候，就会有要求新的两边一定都是要大于当下的两个解的，所以放心，走那个都不会漏解；
//如果后面有大的值 那必然2个都是大于现在的 所以其实2个都移动都行，但是也就快了30ms

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxA := (r - l) * min(height[l], height[r])
	//怎么移动呢那肯定是小的动嘛
	for l < r {
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		maxA = max(maxA, (r-l)*min(height[l], height[r]))
	}
	return maxA
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
