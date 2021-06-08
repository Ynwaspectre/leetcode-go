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
//滑动窗口  这种困难的题有时候就是不敢下手 其实思路都还好 就是有的时候怕自己想错 还是要多做
//这题大概看了下就开始写了 思路和题解的代码差不多了 在处理最后一些小问题的时候提交了多次、
//l和r同时为0 r向右 如果满足则l向右 r不动  不停缩小距离 直到 l和r不满足后 l退1 此时l和r是满足的 记录下最小值和当前索引，如果后面有最小的就更新
//返回最后最小的 其实也还好 要额外写个是否满足的判断方法 ，就是判断数量是否达标 >=为达标 只要<就直接return false
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	//先把t的字母存hash里
	tMap := make(map[uint8]int)
	sMap := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	l, r := 0, 0
	curL := -1
	curR := -1
	minRes := len(s)
	for ; r < len(s); r++ {
		sMap[s[r]]++
		if isManzu(tMap, sMap) {
			//动l 直到不满足
			for isManzu(tMap, sMap) {
				sMap[s[l]]--
				l++
			}
			l--
			sMap[s[l]]++
			if r-l < minRes {
				curL = l
				curR = r
				minRes = r - l
			}
		}
	}
	if curR == -1 {
		return ""
	}
	return s[curL : curR+1] //最后结果是这个
}

func isManzu(t map[uint8]int, s map[uint8]int) bool {
	for k, v := range t {
		if s[k] < v {
			return false
		}
	}
	return true
}
