package 字符串

import "strings"

/**
给你一个二进制字符串 s ，该字符串 不含前导零 。

如果 s 最多包含 一个由连续的 '1' 组成的字段 ，返回 true​​​ 。否则，返回 false 。



示例 1：

输入：s = "1001"
输出：false
解释：字符串中的 1 没有形成一个连续字段。
示例 2：

输入：s = "110"
输出：true


提示：

1 <= s.length <= 100
s[i]​​​​ 为 '0' 或 '1'
s[0] 为 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//最多包含1个由连续1组成的字段  就是1必须是连续的 且遇到0后不能在遇到1了 必然是1xn + 0xm 的格式
//所以只要字符串最后不出现01 这个子串就ok 实在是妙, 自己刚开始还在那循环 不过也解决了
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
