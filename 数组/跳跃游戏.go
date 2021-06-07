package 数组

/**
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。



示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//第一想法是回溯 递归啥的 好像好多要考虑的
//第二想法动态规划从后往前?  不得了 还真是 竟然让我一次写对了 我可真牛 可是为啥时间挺多 按道理是O（n）啊
//等等这好像不是动态规划 因为2者之间没啥联系 i和i-1
func canJump(nums []int) bool {
	length := len(nums)
	df := make([]bool, length)
	df[length-1] = true //最后一个肯定是true

	lastCan := length - 1 //记录最近一个可以的索引
	for i := length - 2; i >= 0; i-- {
		if nums[i] == 0 {
			df[i] = false
		} else {
			df[i] = (i + nums[i]) >= lastCan //大于就是true
			if df[i] {
				lastCan = i
			}
		}
	}
	return df[0]
}
