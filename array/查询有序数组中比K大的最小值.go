package array

//leetcode没找到
//有序数组中 从小到大 找到比给定的K值大的最小值 O(n)比较简单 二分看看
//找到k则返回k的索引 否则返回-1
func findGtKMax(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r + 1) / 2 //取高位
		if nums[mid] <= k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

}
