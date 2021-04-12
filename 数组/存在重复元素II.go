package 数组

/**
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。



示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//利用hash表 存储上一个这个数的索引 如果后面出现和这个索引差值小于等于k的索引 那直接返回false，否则丢到
//目前的索引 换上新的索引 时间大幅度提高  缺点 内存高了
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			//当前索引减去之前的索引如果大于k 就把之前的丢了
			if i-v > k {
				m[nums[i]] = i
			} else {
				return true
			}
		} else {
			m[nums[i]] = i
		}
	}
	return false
}

/**
暴力解决  时间超10%
func containsNearbyDuplicate(nums []int, k int) bool {
	min := 0
	for i := 0; i < len(nums); i++ {
		if (i + k + 1) < len(nums) {
			min = i + k + 1
		} else {
			min = len(nums)
		}
		for j := i + 1; j < min; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
*/
