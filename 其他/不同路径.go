package 其他

/**
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？



示例 1：


输入：m = 3, n = 7
输出：28
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6


提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//递归吧  向下一步就是 m-1,n,向右一步就是 m,n-1
//然鹅超时

//动态规划吧 //去第i行第j列的走法一共有dp[i][j]种 只可能由上2种演变而来
////dp[i][j]=dp[i-1][j]+dp[i][j-1]
//nice 算是第一次不参考任何东西写动态规划 就是在m和n谁是行和列的上面模糊了会,其他思路挺清晰

func uniquePaths(m int, n int) int {
	var dp [101][101]int
	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

//func uniquePaths(m int, n int) int {
//	if m == 1 || n == 1 {
//		return 1
//	}
//	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
//}
