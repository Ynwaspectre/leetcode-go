package 数组

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		s = maxstr(s, strs[i])
		if s == "" {
			return ""
		}
	}
	return s
}

func maxstr(x, y string) string {
	l := min(len(x), len(y))
	s := ""
	for i := 0; i < l; i++ {
		if x[i] == y[i] {
			s += string(x[i])
		}
	}
	return s
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
