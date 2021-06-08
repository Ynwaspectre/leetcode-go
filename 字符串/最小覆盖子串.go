package 字符串

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。



示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"


提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成


进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//你好 能做出来就很吊了 还设计nm的on啊
//看了几分钟 没啥思路 t的顺序不影响结果
//滑动窗口 从刚开始 0
func minWindow(s string, t string) string {
	//先把t的字母存hash里
	tMap := make(map[uint8]int)
	sMap := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	l, r := 0, 0
	for l <= r && r < len(s) {
		//先看是否满足
		sMap[s[r]]++
		if isManzu(tMap, sMap) {
			//满足 移动l
			l++
		} else {
			r++
		}

	}
	return s[l : r+1]
}

func isManzu(t map[uint8]int, s map[uint8]int) bool {
	for k, v := range t {
		if s[k] < v {
			return false
		}
	}
	return true
}
