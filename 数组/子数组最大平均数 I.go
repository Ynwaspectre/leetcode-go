package 数组

import "math"

/**
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。



示例：

输入：[1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75


提示：

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-average-subarray-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//一次ac 思路从头开始遍历 先取最开始k个元素记录sum与平均值，然后一次往后遍历
//去除第一个值加上后面一个值重新计算平均是 并与当前的最大进行比较 最后返回最大的

//看了官方题解 其实不用计算平均数  因为k是固定的 所以只要计算和就ok，但是无伤大雅
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAverage := float64(sum) / float64(k)
	for i := 1; i <= len(nums)-k; i++ {
		sum = sum - nums[i-1] + nums[i+k-1]
		maxAverage = math.Max(maxAverage, float64(sum)/float64(k))
	}
	return maxAverage
}
