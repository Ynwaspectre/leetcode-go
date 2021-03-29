package 数组

//--------------------题目------------------------------
/**
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。


示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 103
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
//-----------------------------------思路--------
/**
构建map
遍历nums 计算target-nums[i]如果这个值存在map里则取出来,否则将当前nums[i]放入map里
即 map[nums[i]]=i
时间复杂度 O(n)
空间复杂度 O(n)
*/
//2020-03-22

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if key, ok := m[target-v]; ok {
			return []int{i, key}
		} else {
			m[v] = i
		}
	}
	return []int{}
}
