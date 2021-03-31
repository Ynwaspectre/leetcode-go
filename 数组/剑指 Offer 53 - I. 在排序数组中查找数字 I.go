package 数组

/**
统计一个数字在排序数组中出现的次数。



示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: 0


限制：

0 <= 数组长度 <= 50000



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//二分吧 一次ac 不过在本地调试了一会 主要思路是找到中间数，然后中间数往两边遍历个数
func search(nums []int, target int) int {
	count := 0
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2 //这边可以优化 mid:=l+(r-l)/2 防止溢出
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			mid1 := mid
			for mid >= 0 && nums[mid] == target {
				count++
				mid--
			}
			for mid1 < len(nums)-1 && nums[mid1+1] == target {
				count++
				mid1++
			}
			break
		}
	}
	return count
}
