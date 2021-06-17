package 二叉树

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。



示例 1：


输入：n = 3
输出：5
示例 2：

输入：n = 1
输出：1


提示：

1 <= n <= 19

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//dp[i] 表示 i个节点的总数
//f[i]表示以i为头结点的总数 所以dp[i]=f[i]+f[i+1]+f[i+2]。。。。+f[n]
//f[i]=dp[i-1]*dp[n-i]  i两边的长度
//f[n]=dp[i-n]*dp[0]
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ { //计算i的 个数 叠加
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

//递归 超出时间限制
//func numTrees(n int) int {
//	if n == 1 || n == 0 {
//		return 1
//	}
//	sum := 0
//	for i := 1; i <= n; i++ {
//		sum += numTrees(i-1) * numTrees(n-i)
//	}
//	return sum
//}
