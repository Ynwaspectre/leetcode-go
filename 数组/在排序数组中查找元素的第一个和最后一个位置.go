package 数组

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？


示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//二分  发现时间差不多哎
func searchRange(nums []int, target int) []int {
	//先找左边  再找右边
	l, r := 0, len(nums)-1
	leftIndex := -1
	rightIndex := -1

	if r == -1 {
		return []int{leftIndex, rightIndex}
	}

	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		leftIndex = l
	}

	l, r = 0, len(nums)-1
	for l < r {
		mid := (l + r + 1) / 2
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if nums[r] == target {
		rightIndex = r
	}

	return []int{leftIndex, rightIndex}
}

////二分查找到任何一个 然后在那个位置往前 往后找 但是复杂度好像不对
//func searchRange(nums []int, target int) []int {
//	l, r := 0, len(nums)-1
//	for l <= r {
//		mid := (l + r + 1) / 2
//		if nums[mid] == target {
//			//找到了 往前找第一个 往后找第二个
//			start := mid
//			for start >= 0 && nums[start] == target {
//				start--
//			}
//			end := mid
//			for end <= len(nums)-1 && nums[end] == target {
//				end++
//			}
//			return []int{start + 1, end - 1}
//		} else if nums[mid] < target {
//			l = mid + 1
//		} else {
//			r = mid - 1
//		}
//	}
//	return []int{-1, -1}
//}
