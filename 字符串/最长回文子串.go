package 字符串

/**
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//没思路没思路 唯一思路就是暴力判断。。。先看会别的吧
//双指针吧 固定一个在左边 右边的从1开始 但是感觉还是很烦
//不小心点开提示 发现有动态规划，之前思考的时候确实有这么点，为什么呢
//因为存在包含关系，有的可能会重复 那么方程和初始化呢

//dp[i][j] 表示第字符串s第i到第j的子串是否是回文
//粗略看了题解 这转移方程我应该想不出来 但是我就大概粗略看了下 断断续续写出来了
//思路懂了 问题在于如何循环，因为i这个变量需要i+1来支持 所以从后往前 j没事

//如果i==j dp[i][j]肯定是true
//如果i+1==j 则dp[i][j]=s[i]==s[j]
//其余是 dp[i][j]=dp[i+1][j-1]&&s[i]==s[j]

//时间有点慢。。

func longestPalindrome(s string) string {
	var dp [1001][1001]bool
	for i := len(s) - 2; i >= 0; i-- {
		for j := len(s) - 1; j >= i; j-- {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
		}
	}
	max := 0
	maxStr := string(s[0])
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if dp[i][j] {
				if j-i+1 > max {
					max = j - i + 1
					maxStr = s[i : j+1]
				}
			}
		}
	}
	return maxStr
}
