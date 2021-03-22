package array

import (
	"fmt"
)

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

2. 不使用除法 可以就着题目的 来想 暴力点

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
输入: [1,2,3,4,5]
输出: [120,60,40,30,24]
3.
*/
func constructArr(a []int) []int {
	b := make([]int, len(a))
	b[0] = 1
	for i := 0; i < len(a); i++ {
		b[i] *= a[i]
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Print(constructArr(nums))
}
