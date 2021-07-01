package 数组

/**
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//dp[i]代表 0~i的字符串是否能被wordDict匹配到
//i从1开始 先默认dp[i]=true 不然后面的值都是false 不好计算
//dp[i]=dp[i-1]&&s[i] 点j在i0到i-1上  dp[i]满足任何一个即可
//dp[i]=(dp[i-1]&&s[i])||dp[i-2]&&s[i-2:i]||.....-3,-4
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if !dp[j] {
				continue
			}
			if check(s[j:i], wordDict) {
				dp[i] = true
				continue
			}

		}
	}
	return dp[len(s)]
}

func check(s string, wordDict []string) bool {
	for i := 0; i < len(wordDict); i++ {
		if s == wordDict[i] {
			return true
		}
	}
	return false
}

// 写了个超时的dfs。。
//func wordBreak(s string, wordDict []string) bool {
//	if s == "" {
//		return false
//	}
//	result := false
//	//找到了前缀就可以缩短s的长度了 查看s是否以某个字符串开头
//	for i := 0; i < len(wordDict); i++ {
//		if startWith(s, wordDict[i]) {
//			if len(s) == len(wordDict[i]) {
//				return true
//			}
//			result = result || wordBreak(s[len(wordDict[i]):], wordDict)
//		}
//	}
//	return result
//}
//
//func startWith(s, sub string) bool {
//	l := len(sub)
//	if l > len(s) {
//		return false
//	}
//	return s[:l] == sub
//}
