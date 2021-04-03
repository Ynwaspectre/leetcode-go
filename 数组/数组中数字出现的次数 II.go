package 数组

import "sort"

/**

在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。



示例 1：

输入：nums = [3,4,3,3]
输出：4
示例 2：

输入：nums = [9,1,7,9,7,9,7]
输出：1


限制：

1 <= nums.length <= 10000
1 <= nums[i] < 2^31

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//比较暴力。。直接拍完序 然后比较3个数字的首尾 如果相同则3个3个的递增i,否则不同数字就是i,
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i = i + 3 {
		if nums[i] != nums[i+2] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
