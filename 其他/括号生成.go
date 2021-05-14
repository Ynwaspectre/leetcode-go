package 其他

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]


提示：

1 <= n <= 8
通过次数272,949提交次数353,854

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//暴力生成所有 在判断是否有效？"("
//第一个肯定是左括号 最右边肯定是 ")"
//中间就是 n-1的结果呀
//这个思路ok的 做了一会儿也完成了 而且时间几百100% 内存有点多3.2MB 问题不大
//一次解决
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	if n == 2 {
		return []string{"()()", "(())"}
	}
	//合并去重
	a := generateParenthesis(n - 1) //n-1的结果 然后用一对括号在每个位置插入生成一个结果
	var result []string
	resultMap := make(map[string]int)
	for _, v := range a {
		for i := 0; i < len(v); i++ {
			if _, ok := resultMap[v[:i]+"()"+v[i:]]; !ok {
				result = append(result, v[:i]+"()"+v[i:])
				resultMap[v[:i]+"()"+v[i:]]++
			}
		}
	}
	return result
}
