package 数组

/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。



示例 1：

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
示例 2：

输入：nums = [0]
输出：[[],[0]]


提示：

1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//子集  思路和之前的全排列好像差不多 我这个 就是用新的元素与现有里面的每个在组合一个
func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {

		j := 0
		length := len(result)
		for j < length {
			copyResultJ := make([]int, len(result[j])) //这一行要复制下  哎 对切片的理解还是不够导致在切片append的时候经常容易出错
			copy(copyResultJ, result[j])
			result = append(result, append(copyResultJ, nums[i]))
			j++
		}
	}
	return result
}
