package 其他

/**
请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//单调栈  一个单调递减栈  受到之前题目的影响
func dailyTemperatures(temperatures []int) []int {
	var stack []int
	l := len(temperatures)
	index := make([]int, len(temperatures))
	for i := 0; i < l; i++ {
		if len(stack) == 0 {
			stack = append(stack, 0)
		} else {
			//大于栈顶最小的 出栈记录数据
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
				index[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	//全部出栈 设为0
	for len(stack) > 0 {
		index[stack[len(stack)-1]] = 0
		stack = stack[:len(stack)-1]
	}
	return index
}
