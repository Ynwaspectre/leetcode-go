package 字符串

import "strings"

/**
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同



示例 1：

输入：s = "bcabc"
输出："abc"
示例 2：

输入：s = "cbacdcbc"
输出："acdb"


提示：

1 <= s.length <= 104
s 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicate-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeDuplicateLetters(s string) string {
	m := make(map[string]int)
	var stack []string
	exist := make(map[string]int)
	//统计次数
	for _, v := range s {
		m[string(v)]++
	}
	//栈里面大于当前元素的都出来
	stack = append(stack, string(s[0]))
	exist[string(s[0])]++
	for i := 1; i < len(s); i++ {
		if _, ok := exist[string(s[i])]; ok {
			//存在就不放进去了
			m[string(s[i])]-- //数量-1
			continue
		}
		//把大于当前元素值弄出来  这样可以获得当前的最小
		for len(stack) > 0 && stack[len(stack)-1] > string(s[i]) && m[stack[len(stack)-1]] > 1 {
			m[stack[len(stack)-1]]--
			delete(exist, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, string(s[i]))
		exist[string(s[i])]++
	}
	return strings.Join(stack, "")
}
