package array

/*
Q:
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。



示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]


提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

/**
1.	如果可以使用除法的话 可以遍历数组取的整个数组所有数的积
	然后再次遍历数组 使用该积除以每个数
	时间复杂度 O(n)
	空间复杂度 O(1)

2. 不使用除法 可以就着题目的 来想 暴力点 果断超出时间限制

func baoli(a []int) []int {
	b:=make([]int,len(a))
	for k, _ := range a {
		b[k]=baolijisuan(a,k)
	}
	return b
}

func baolijisuan(a []int, key int) int {
	ji:=1
	for i := 0; i < len(a); i++ {
		if i==key {
			continue
		}
		ji*=a[i]
	}
	return ji
}
输入: [1,2,3,4,5] a
输出: [120,60,40,30,24] b
3.
对于b[i]  以上述输入输出为例 先看b[1],
@1 第一次遍历
b[i]的值 等于a[0]~a[i-1]的积 * a[i+1]~a[len(a)-1]的积
比如 b中的60就等于 左边的积1*右边的积3*4*5
@ 第二次遍历
按照上面的思路即需要倒序遍历
取个中间变量temp=1
b[i]=temp*b[i+1]


*/
func constructArr(a []int) []int {
	if len(a) == 0 { //这边不能忘记加 不然ac不掉 空数组返回空数组
		return a
	}
	b := make([]int, len(a)) //构建和a长度一样的b切片
	b[0] = 1                 //b[i] = b[i-1] * a[i] 当i=1的时候b[0]没值 所以取b[0]=1 因为第一个元素是
	///最左边，在左边没有值了 积就是1
	//对左边遍历
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	//右边遍历 这儿比较难想 参考上面的想法倒过来
	temp := 1
	for i := len(a) - 2; i >= -1; i-- {
		b[i+1] *= temp //两次的值相乘 得最后的
		temp *= a[i+1]
	}
	return b
}
