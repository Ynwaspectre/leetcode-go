package 数组

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。



示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]


提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//nums的全排列等于nums[:len(nums)-1]的全排列 与最后一个数组合
//递归，但是在切片处理那边有点问题 调试了一会 对切片的理解还是不够深
func permute(nums []int) [][]int {
	length := len(nums)
	if length == 1 {
		return [][]int{nums}
	}
	if length == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	result := permute(nums[:length-1])
	var newResult [][]int
	var lastElement = nums[length-1]
	for _, v := range result {
		for i := 0; i <= len(v); i++ {
			copyPrefix := make([]int, i)
			copy(copyPrefix, v[:i])
			k := append(copyPrefix, lastElement)
			copyEndfix := make([]int, len(v)-i)
			copy(copyEndfix, v[i:])
			newElement := append(k, copyEndfix...)
			newResult = append(newResult, newElement)
		}
	}
	return newResult
}
