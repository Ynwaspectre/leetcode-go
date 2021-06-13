package 其他

import "strings"

/**
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :

输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-k-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//想了好久想出来了 也是在知道要用单调栈的基础上。。和官方思路一样 还算ok吧 自己想出来的应该
//比看题解稍微好点
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	var stack []string
	stack = append(stack, string(num[0]))
	for i := 1; i < len(num); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > string(num[i]) && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, string(num[i]))
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	i := 0
	for stack[i] == "0" && len(stack) > 1 {
		stack = stack[i+1:]
	}
	return strings.Join(stack, "")
}
