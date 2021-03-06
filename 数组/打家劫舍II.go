package 数组

/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。

给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。



示例 1：

输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2：

输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 3：

输入：nums = [0]
输出：0


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//哈哈 这个是我做了之前系列实打实想出来的 不过也有之前看题解模糊印象的功劳
//就是分成了2个数组 来做各自的存储 然后取最大值
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	var dp [500]int
	var dp1 [500]int
	aNums := nums[:len(nums)-1]
	bNums := nums[1:]

	dp[0] = aNums[0]
	dp[1] = max(aNums[0], aNums[1])
	for i := 2; i < len(aNums); i++ {
		dp[i] = max(dp[i-2]+aNums[i], dp[i-1])
	}

	dp1[0] = bNums[0]
	dp1[1] = max(bNums[0], bNums[1])
	for i := 2; i < len(bNums); i++ {
		dp1[i] = max(dp1[i-2]+bNums[i], dp1[i-1])
	}
	return max(dp[len(aNums)-1], dp1[len(bNums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
