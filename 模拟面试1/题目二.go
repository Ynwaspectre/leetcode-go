package 模拟面试1

/**
给你一个整数 n ，请你找出并返回第 n 个 丑数 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

输入：n = 10
输出：12
解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
示例 2：

输入：n = 1
输出：1
解释：1 通常被视为丑数。

5 4 6 8 10

提示：

1 <= n <= 1690
*/
//好像没啥思路 这题表述的太拗口了
//n==10 则依次输出10个丑数  然后返回最后一个丑数

//看的题解 这谁想的出来啊。。。太牛
// 1, 2, 3, 4, 5, 6, 8, 9, 10, 12
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1 //第一个丑数是1
	//第二个是 3*dp[i] 2*dp[i] 5*dp[i]的最小值
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//循环的方法超时
//func nthUglyNumber(n int) int {
//	if n <= 6 {
//		return n
//	}
//	n = n - 6
//	for i := 7; i < math.MaxInt64; i++ {
//		j:=i
//		for j%2 == 0 {
//			j = j / 2
//		}
//		for j%3 == 0 {
//			j= j / 3
//		}
//		for j%5 == 0 {
//			j = j / 5
//		}
//		if j == 1 {
//			n--
//		}
//		if n == 0 {
//			return i
//		}
//	}
//	return -1
//}
