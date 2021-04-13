package 字符串

/**

给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/
//没想出来 为什么么呢 第一时间想到是判断遇到2个不一样的 但是在想的时候我是在想要删哪个然后继续往下走
//然后统计删除的个数 如果后面还要删除 则是false 后面不用删除了则是true
//因为遇到2个不同的肯定要删一个 所以分2种情况 先删左边的 看右边剩下的是否是回文字符串 结果是res1  删右边的看左边的
//是否是回文字符串结果是res2的话 那么只要res1和res2有一个是回文就会满足 所以return res1||res2
func validPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			//则肯定会删除其中一个 如果删除l
			return isValidPalindome(s[l+1:r+1]) || isValidPalindome(s[l:r])
		}
		l++
		r--
	}
	return true
}

func isValidPalindome(s string) bool {
	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
