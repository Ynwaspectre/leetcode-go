package 数组

import (
	"sort"
)

/**
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// [-1,-2,0,1,2,3,4]
//从早上到下午一直在弄，思路和题解一样  去重本来用的是map的 有点耗时 根据题解改了下好了点
func threeSum(nums []int) [][]int {
	var result [][]int //结果
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	i, j, k := 0, 1, length-1
	for i < length-1 && nums[i] <= 0 {
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				//找到后如何移动
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			} else if sum < 0 {
				//小于0的时候该如何移动
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			} else {
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			}
		}
		for i < length-1 && nums[i] == nums[i+1] {
			i++
		}
		i++
		j = i + 1
		k = length - 1
	}
	return result
}
