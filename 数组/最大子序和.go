package 数组

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [0]
输出：0
示例 4：

输入：nums = [-1]
输出：-1
示例 5：

输入：nums = [-100000]
输出：-100000


提示：

1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105


进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//最大的肯定2边都是负数 这种感觉一看就是要双指针
func maxSubArray(nums []int) int {
	l := 0
	r := len(nums) - 1
	sum := 0
	for _, v := range nums {
		sum += v
	}
	sum1 := sum
	for l < r {
		if nums[l] <= nums[r] {
			l++
			sum1 -= nums[l]
		} else {
			r--
			sum1 -= nums[r]
		}
		sum = max(sum1, sum)
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
