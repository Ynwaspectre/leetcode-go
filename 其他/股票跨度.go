package 其他

/**
编写一个 StockSpanner 类，它收集某些股票的每日报价，并返回该股票当日价格的跨度。

今天股票价格的跨度被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。

例如，如果未来7天股票的价格是 [100, 80, 60, 70, 60, 75, 85]，那么股票跨度将是 [1, 1, 1, 2, 1, 4, 6]。



示例：

输入：["StockSpanner","next","next","next","next","next","next","next"], [[],[100],[80],[60],[70],[60],[75],[85]]
输出：[null,1,1,1,2,1,4,6]
解释：
首先，初始化 S = StockSpanner()，然后：
S.next(100) 被调用并返回 1，
S.next(80) 被调用并返回 1，
S.next(60) 被调用并返回 1，
S.next(70) 被调用并返回 2，
S.next(60) 被调用并返回 1，
S.next(75) 被调用并返回 4，
S.next(85) 被调用并返回 6。

注意 (例如) S.next(75) 返回 4，因为截至今天的最后 4 个价格
(包括今天的价格 75) 小于或等于今天的价格。


提示：

调用 StockSpanner.next(int price) 时，将有 1 <= price <= 10^5。
每个测试用例最多可以调用  10000 次 StockSpanner.next。
在所有测试用例中，最多调用 150000 次 StockSpanner.next。
此问题的总时间限制减少了 50%。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/online-stock-span
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//维护一个单调递增栈 小于当前元素的都移除栈
type StockSpanner struct {
	Stack  []int //存price
	Weight []int //记录每个price距离前面第一个比他大的数的距离
}

func Constructor() StockSpanner {
	spanner := StockSpanner{[]int{}, []int{}}
	return spanner
}

func (this *StockSpanner) Next(price int) int {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, price)
		this.Weight = append(this.Weight, 1)
		return 1
	} else {
		w := 1
		for len(this.Stack) > 0 && price >= this.Stack[len(this.Stack)-1] {
			this.Stack = this.Stack[:len(this.Stack)-1] //移除出栈
			w += this.Weight[len(this.Weight)-1]
			this.Weight = this.Weight[:len(this.Weight)-1]
		}
		this.Stack = append(this.Stack, price)
		this.Weight = append(this.Weight, w)
		return w
	}
}

//算是个暴力。。 每次都要和前面的比较
//type StockSpanner struct {
//	Stack  []int
//}
//
//func Constructor() StockSpanner {
//	spanner := StockSpanner{[]int{}}
//	return spanner
//}
//
//func (this *StockSpanner) Next(price int) int {
//	if len(this.Stack) == 0 {
//		this.Stack = append(this.Stack, price)
//		return 1
//	} else {
//		count := 1
//		j:=len(this.Stack)-1
//		for j >= 0 && price >= this.Stack[j] {
//			count++
//			j--
//		}
//		this.Stack = append(this.Stack, price)
//		return count
//	}
//}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
