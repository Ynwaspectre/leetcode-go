package 数组

/**
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//排个序先！双指针但是好像不咋好实现 还要考虑相同元素多个情况，咋整
// 没思路 //
//终于在看了题解后 理解了会写出来了 之前没遇到过这种 算是一种新的见识了，
//其实我刚开始想到了递归 但是方式不一样  还有第60行和61行 append要append复制的一份 没搞懂
func combinationSum(candidates []int, target int) [][]int {
	var combine []int
	result := dfs(candidates, target, combine, 0)
	return result
}

//
func dfs(candidates []int, target int, combine []int, idx int) [][]int {
	var result [][]int
	if idx == len(candidates) {
		return result
	}
	if target == 0 {
		newCombine := make([]int, len(combine))
		copy(newCombine, combine)
		result = append(result, newCombine)
		return result
	}
	result = append(result, dfs(candidates, target, combine, idx+1)...)
	if target-candidates[idx] >= 0 {
		combine = append(combine, candidates[idx])
		result = append(result, dfs(candidates, target-candidates[idx], combine, idx)...)
	}
	return result
}
