package 其他

/**
请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]


限制：

1 <= push_back,pop_front,max_value的总操作数 <= 10000
1 <= value <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//这题刚开始想简单了,最大值max只用了一个值来存储的,考虑到题目说的O1的操作,我还用的链表
//后面改成了max用一个切片来维护, 而且在处理弹出队列和max值的时候就感觉很简单 但是写起来就是O（n）对于
//均摊复杂度也不是很了解,处理max可以用单调栈来处理,不一定要每个元素都放进去，只需要把小于我的元素踢出来，然后把自己放进去,如果没有小于
//自己的元素 可以不放进去
//链表也可以换成另外一个切片

//核心其实还是处理max,这题的关键
//假设5个数  3 4 5 2 1   从头到后的顺序放的  最后按3 4 5 2 1 的顺序出来

//-----进入
//max 第一次放3  max = [3]
//max 第二次放4  max =[4] //把比4小的数都去掉 因为之前的数比4小的出去的时候最大值都是4 那些数后期就没有任何参考价值了 那些数的max只要记录当前的最大值就ok
//max 第三次放5  max=[5] //同上
//max 第四次放2  max=[2,5]//这个2 要保留 因为2是在5后面才出去 所以5出去后肯定要有一个记录目前的最大值
//max 第五次放1 max=[1,2,5]//同上

//-----此时的max就是 [1,2,5]

//出去
// 第一个出3 结果[4,5,2,1]   最大值是5 不动
//第二个出4  [5,2,1] 最大值是5 不动
//第三个出5  [2,1] max 的5 相应的出了 max=[1,2]
//第四个出2 [1] max的2 也出去 max=[1]
//第五个出1 []  max的1 也出去 max=[]
//over

type MaxQueue struct {
	maxArr []int //装max的数组 最大值放在最后面
	resArr []int //主队列
}

func Constructor() MaxQueue {
	var maxArr []int
	var resArr []int
	return MaxQueue{maxArr, resArr}
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxArr) == 0 {
		return -1
	}
	return this.maxArr[len(this.maxArr)-1]
}

func (this *MaxQueue) Push_back(value int) {
	this.resArr = append(this.resArr, value)
	//从前往后 如果有小于value的值都丢掉 没有小于value的 value要放进去
	for len(this.maxArr) > 0 && value > this.maxArr[0] {
		this.maxArr = this.maxArr[1:]
	}
	this.maxArr = append([]int{value}, this.maxArr...)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.maxArr) == 0 {
		return -1
	}
	result := this.resArr[0]
	this.resArr = this.resArr[1:]
	if result == this.Max_value() {
		this.maxArr = this.maxArr[:len(this.maxArr)-1]
	}
	return result
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
