package 字符串

/**
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//观察了下 感觉就是找2个字符串的最大公共子串 word1可以中间有间隔 毕竟可以删除
//word1的最大长度是 word1[:len(word1)-1]的最大长度 选或者不选 最后一个
//如何word2包含这个 就可以成立  事实做不出来-_-||

//看了官方的大概思路
/**  word1=horse   word2=ros
//dp[i][j] 代表word1的前i个字符串编辑到word2的前j个字符串的编辑距离
//  此值等于
		dp[i-1][j]+1    word1的前i-1个字符串到 word2的前j个字符串的编辑距离+1,只要增加一个word2[j]  hors=>ros  只要最后一步删除最后一个e
		dp[i][j-1]+1    word1的前i个字符串到 word2的前j-1个字符串的编辑距离+1  只要word1 删除一个  horse=>ro   只需要最后一步加一个s
		dp[i-1][j-1] 或者 dp[i-1][j-1]+1   取决于 word1和word2最后一个字符是否相同 不相同就+1 因为需要多一次替换操作 hors=》ro   如果e==s 则不需要修改 否则需要修改+1不
*/

//dp[i][j] 代表word1的前i个字符串编辑到word2的前j个字符串的编辑距离
//horse ros  h r  ho   r
func minDistance(word1 string, word2 string) int {
	dp := [501][501]int{}
	m := len(word1)
	n := len(word2)
	//h=>""  ho=>""   word2是空字符串 所以次数就是word1的长度
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	//""=>r  ""=>ro ""=>ros 同理 word1是空字符串 次数是word2的长度
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1)
			if word1[i-1] == word2[j-1] { //这儿是长度-1的索引
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
