package 数组

import "sort"

/**
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。



示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
示例 2：

输入：nums = [1,2,3,4]
输出：0
示例 3：

输入：nums = [1]
输出：0


提示：

1 <= nums.length <= 104
-105 <= nums[i] <= 105


进阶：你可以设计一个时间复杂度为 O(n) 的解决方案吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//排序  从前开始找第一个不一样的  从后往前 找第一个不一样的
func findUnsortedSubarray(nums []int) int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)
	l := 0
	r := 0
	for i := 0; i < len(copyNums); i++ {
		if copyNums[i] != nums[i] {
			l = i
			break
		}
	}
	for i := len(copyNums) - 1; i >= 0; i-- {
		if copyNums[i] != nums[i] {
			r = i
			break
		}
	}
	if r == 0 && l == 0 {
		return 0
	}
	return r - l + 1
}
