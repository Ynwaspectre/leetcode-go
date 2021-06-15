package 模拟面试1

/**
有台奇怪的打印机有以下两个特殊要求：

打印机每次只能打印由 同一个字符 组成的序列。
每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
给你一个字符串 s ，你的任务是计算这个打印机打印它需要的最少打印次数。


示例 1：

输入：s = "aaabbb"
输出：2
解释：首先打印 "aaa" 然后打印 "bbb"。
示例 2：

输入：s = "aba"
输出：2
解释：首先打印 "aaa" 然后在第二个位置打印 "b" 覆盖掉原来的字符 'a'。


提示：

1 <= s.length <= 100
s 由小写英文字母组成
*/

//

/**
直接进行一个不会 想到动态规划了 但是没想到用二维
左边和右边相等有特殊性也想到了 可惜细想不下去
看了题解自己写还是能完成80% 左右的
dp[i][j] 表示打印字符的i位置到j位置
abacd
如果 s[i]==s[j] 那打印室s[i]的时候肯定会顺便把s[j]打印了  dp[i][j]= dp[i][j-1]
如果 s[i]!=s[j]     那就是打印s[i][k] + s[k+1][j] 找到一个k满足
*/
func strangePrinter(s string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		m := make([]int, len(s)+1)
		dp[i] = m
	}
	for i := len(s) - 1; i >= 0; i-- { //这边i倒序 我还是可以理解的 因为下面有计算k+1的 如果i不是倒叙 k+1的就没结果
		dp[i][i] = 1                      //这个也可以理解 本来我还傻傻的 在 j循环下面 判断i==j 然后赋值 dp[i][j]=1 问题不大
		for j := i + 1; j < len(s); j++ { //j升序有j-1 所以也算可以理解 但是j为什么要赋值i+1 哦后面想了下 因为我们默认i~j i在前面 所以j从i+1开始 当i等于len(s-2)的时候
			// 其实就是打印最后2个字符
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				minCount := j - i + 1
				for k := i; k < j; k++ {
					temp := dp[i][k] + dp[k+1][j]
					minCount = min(minCount, temp)
				}
				dp[i][j] = minCount
			}
		}
	}
	return dp[0][len(s)-1]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
