package 数组

/**
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。



进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


提示：

0 <= nums.length <= 104
-109 <= nums[i] <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//hash
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	maxValue := 1
	for k, _ := range m {
		if _, ok := m[k-1]; ok {
			continue
		}
		j := k + 1
		count := 1
		_, ok := m[j]
		for ok {
			j++
			_, ok = m[j]
			count++
		}
		maxValue = max(maxValue, count)
	}
	return maxValue
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//排序完遍历一下  但是复杂度是 n*logN
//func longestConsecutive(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	sort.Ints(nums)
//	count := 1
//	maxValue := 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i]-nums[i-1] == 1 {
//			count++
//			maxValue = max(maxValue, count)
//		} else if nums[i]-nums[i-1] == 0 {
//			continue
//		} else {
//			count = 1
//		}
//	}
//	return maxValue
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
